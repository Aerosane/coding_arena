#!/usr/bin/env python3
"""
Local sanity check for the judge test solutions.

Layout: <problem>/<ac|tle>/{main.cpp, main.py, Main.java}

Phase 1 - official tests (../problems/<problem>/*.in|*.out)
    ac  solutions must be Accepted.
    tle solutions must ALSO be correct here (the official tests are tiny, so a
        slow-but-correct algorithm finishes instantly and is accepted).
Phase 2 - generated stress inputs (nothing is written to judge-config/problems)
    ac  solutions must still be Accepted.
    tle solutions must exceed the time limit.

Stress sizes are intentionally LARGER than the official constraints in some
cases (see STRESS below), because with the official limits an O(N^2) solution
cannot reliably TLE in C++/Java (e.g. two-sum N<=1e4, merge m+n<=200). These
inputs only validate that the slow solutions really are slow; do not add the
oversized ones to the judge unless the problem constraints are raised too.

Usage:
    python judge-config/solutions/run_solutions.py [problem ...] [--no-stress]
    TL=2 python ...      # wall-clock limit in seconds (default 2; JVM start-up
                         # makes a strict 1s unreliable locally)

Needs g++, python, and a JDK (javac) on PATH.
"""
import os
import shutil
import subprocess
import sys
import tempfile
import time
from pathlib import Path

ROOT = Path(__file__).resolve().parent
PROBLEMS = ROOT.parent / "problems"
TL = float(os.environ.get("TL", "2.0"))
LANGS = ["cpp", "py", "java"]
LABEL = {"cpp": "C++", "py": "Python", "java": "Java"}


# ----------------------------------------------------------------- stress data
def _two_sum():
    n = 200_000  # official max 1e4
    nums = [2 * i for i in range(n - 2)] + [400_001, 400_003]  # only odd+odd pair hits target
    inp = f"{n}\n{' '.join(map(str, nums))}\n{400_001 + 400_003}\n"
    return inp, f"{n - 2} {n - 1}\n"


def _reverse():
    s = "".join(chr(97 + i % 26) for i in range(1_000_000))  # official max 1e5
    return s + "\n", s[::-1] + "\n"


def _parens():
    s = "(" * 200_000 + ")" * 200_000  # official max 1e4; fully nested = worst case
    return s + "\n", "true\n"


def _merge():
    m = n = 100_000  # official max m+n = 200
    a = [2 * i for i in range(m)]
    b = [2 * i + 1 for i in range(n)]
    inp = f"{m} {n}\n{' '.join(map(str, a))}\n{' '.join(map(str, b))}\n"
    return inp, " ".join(map(str, sorted(a + b))) + "\n"


def _longest():
    cyc = [chr(c) for c in range(33, 127)]  # 94 distinct printable chars
    s = "".join(cyc[i % 94] for i in range(200_000))  # official max 5e4
    return s + "\n", "94\n"


STRESS = {
    "two-sum": _two_sum,
    "reverse-string": _reverse,
    "valid-parentheses": _parens,
    "merge-sorted-arrays": _merge,
    "longest-substring": _longest,
}


# --------------------------------------------------------------------- helpers
def build(lang, src_dir, work):
    if lang == "cpp":
        exe = work / "sol"
        subprocess.run(["g++", "-O2", "-std=c++17", "-o", str(exe), str(src_dir / "main.cpp")], check=True)
        return [str(exe)]
    if lang == "py":
        return [sys.executable, str(src_dir / "main.py")]
    javac = shutil.which("javac")
    cmd = [javac] if javac else ["java", "-m", "jdk.compiler/com.sun.tools.javac.Main"]
    subprocess.run(cmd + ["-d", str(work), str(src_dir / "Main.java")], check=True)
    return ["java", "-XX:+UseSerialGC", "-cp", str(work), "Main"]


def norm(text):
    lines = [ln.rstrip() for ln in text.replace("\r\n", "\n").split("\n")]
    while lines and lines[-1] == "":
        lines.pop()
    return lines


def run_case(cmd, stdin_bytes, expected_text):
    start = time.time()
    try:
        p = subprocess.run(cmd, input=stdin_bytes, capture_output=True, timeout=TL)
    except subprocess.TimeoutExpired:
        return "TLE", time.time() - start
    took = time.time() - start
    if p.returncode != 0:
        return "RE", took
    ok = norm(p.stdout.decode(errors="replace")) == norm(expected_text)
    return ("AC" if ok else "WA"), took


def phase_official(cmd, problem_dir):
    slowest = 0.0
    for f in sorted(problem_dir.glob("*.in"), key=lambda p: int(p.stem)):
        v, t = run_case(cmd, f.read_bytes(), f.with_suffix(".out").read_text())
        slowest = max(slowest, t)
        if v != "AC":
            return f"{v} (test {f.stem})", slowest
    return "AC", slowest


def phase_stress(cmd, slug):
    inp, out = STRESS[slug]()
    v, t = run_case(cmd, inp.encode(), out)
    return v, t


def main():
    args = [a for a in sys.argv[1:] if not a.startswith("--")]
    do_stress = "--no-stress" not in sys.argv
    slugs = sorted(d.name for d in ROOT.iterdir() if d.is_dir() and (not args or d.name in args))
    bad = 0
    print(f"{'problem':<21}{'lang':<8}{'kind':<5}{'official':<16}{'stress':<10}result")
    for slug in slugs:
        for kind in ("ac", "tle"):
            for lang in LANGS:
                with tempfile.TemporaryDirectory() as tmp:
                    try:
                        cmd = build(lang, ROOT / slug / kind, Path(tmp))
                    except (subprocess.CalledProcessError, FileNotFoundError) as e:
                        print(f"{slug:<21}{LABEL[lang]:<8}{kind:<5}BUILD FAILED ({e})")
                        bad += 1
                        continue
                    off, t1 = phase_official(cmd, PROBLEMS / slug)
                    ok = off == "AC"  # both kinds must be correct on official tests
                    st = "-"
                    if do_stress:
                        st, t2 = phase_stress(cmd, slug)
                        ok = ok and (st == "AC" if kind == "ac" else st == "TLE")
                        st = f"{st} {t2:.1f}s"
                bad += not ok
                print(f"{slug:<21}{LABEL[lang]:<8}{kind:<5}{off + f' {t1:.2f}s':<16}{st:<10}{'ok' if ok else 'FAIL'}")
    print("\nall solutions behave as expected" if not bad else f"\n{bad} solution(s) did not behave as expected")
    sys.exit(1 if bad else 0)


if __name__ == "__main__":
    main()
