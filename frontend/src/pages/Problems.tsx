import { useEffect, useState, useMemo } from 'react'
import { useNavigate } from 'react-router-dom'
import { getProblems, type Problem } from '@/lib/data-service'
import { ErrorBoundary } from '@/components/ErrorBoundary'

export default function ProblemsPage() {
  const [problems, setProblems] = useState<Problem[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)
  const [difficultyFilter, setDifficultyFilter] = useState<string>('All')
  const [categoryFilter, setCategoryFilter] = useState<string>('All')
  const navigate = useNavigate()

  const loadProblems = async () => {
    try {
      setLoading(true)
      setError(null)
      const data = await getProblems()
      setProblems(data)
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to load problems')
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    loadProblems()
  }, [])

  // Get unique categories from problems
  const categories = useMemo(() => {
    const uniqueCategories = new Set(problems.map(p => p.category))
    return ['All', ...Array.from(uniqueCategories).sort()]
  }, [problems])

  // Filter problems based on selected criteria
  const filteredProblems = useMemo(() => {
    return problems.filter(problem => {
      const matchesDifficulty = difficultyFilter === 'All' || problem.difficulty === difficultyFilter
      const matchesCategory = categoryFilter === 'All' || problem.category === categoryFilter
      return matchesDifficulty && matchesCategory
    })
  }, [problems, difficultyFilter, categoryFilter])

  const handleProblemClick = (problemId: string) => {
    navigate(`/problem/${problemId}`)
  }

  const getDifficultyColor = (difficulty: string) => {
    switch (difficulty) {
      case 'Easy':
        return 'text-green-600'
      case 'Medium':
        return 'text-yellow-600'
      case 'Hard':
        return 'text-red-600'
      default:
        return 'text-gray-600'
    }
  }

  if (loading) {
    return (
      <div className="min-h-screen bg-gray-50 text-gray-800 p-4 sm:p-8">
        <div className="max-w-7xl mx-auto">
          <h1 className="text-2xl sm:text-3xl font-bold mb-6 sm:mb-8 text-gray-900">Practice Problems</h1>
          <div className="flex items-center justify-center py-12">
            <div className="flex flex-col items-center gap-4">
              <div className="w-12 h-12 border-4 border-gray-300 border-t-green-600 rounded-full animate-spin"></div>
              <div className="text-gray-600">Loading problems...</div>
            </div>
          </div>
          {/* Skeleton screens */}
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-5">
            {[1, 2, 3, 4, 5, 6].map((i) => (
              <div
                key={i}
                className="bg-white rounded-lg p-6 border border-gray-200 shadow-sm animate-pulse"
              >
                <div className="h-6 bg-gray-200 rounded w-3/4 mb-4"></div>
                <div className="flex items-center gap-4 mb-4">
                  <div className="h-4 bg-gray-200 rounded w-16"></div>
                  <div className="h-4 bg-gray-200 rounded w-20"></div>
                </div>
                <div className="flex items-center justify-between">
                  <div className="h-4 bg-gray-200 rounded w-16"></div>
                  <div className="h-4 bg-gray-200 rounded w-20"></div>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    )
  }

  if (error) {
    return (
      <div className="min-h-screen bg-gray-50 text-gray-800 p-4 sm:p-8">
        <div className="max-w-7xl mx-auto">
          <h1 className="text-2xl sm:text-3xl font-bold mb-6 sm:mb-8 text-gray-900">Practice Problems</h1>
          <div className="bg-red-50 border border-red-200 rounded-lg p-6 shadow-sm">
            <p className="text-red-700 mb-4">{error}</p>
            <button
              onClick={loadProblems}
              className="px-5 py-2.5 bg-green-600 hover:bg-green-700 text-white font-medium rounded-md transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-green-500"
            >
              Retry
            </button>
          </div>
        </div>
      </div>
    )
  }

  return (
    <ErrorBoundary>
      <div className="min-h-screen bg-gray-50 text-gray-800 p-4 sm:p-8">
        <div className="max-w-7xl mx-auto">
          <h1 className="text-2xl sm:text-3xl font-bold mb-6 sm:mb-8 text-gray-900">Practice Problems</h1>
          
          {/* Filter Controls */}
          <div className="mb-6 flex gap-3 sm:gap-4 flex-wrap items-center">
            <div className="flex items-center gap-2">
              <label htmlFor="difficulty-filter" className="text-gray-700 text-sm font-medium">
                Difficulty:
              </label>
              <select
                id="difficulty-filter"
                value={difficultyFilter}
                onChange={(e) => setDifficultyFilter(e.target.value)}
                className="bg-white border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:border-green-500 focus:ring-2 focus:ring-green-200 transition-all duration-200"
              >
                <option value="All">All</option>
                <option value="Easy">Easy</option>
                <option value="Medium">Medium</option>
                <option value="Hard">Hard</option>
              </select>
            </div>

            <div className="flex items-center gap-2">
              <label htmlFor="category-filter" className="text-gray-700 text-sm font-medium">
                Category:
              </label>
              <select
                id="category-filter"
                value={categoryFilter}
                onChange={(e) => setCategoryFilter(e.target.value)}
                className="bg-white border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:border-green-500 focus:ring-2 focus:ring-green-200 transition-all duration-200"
              >
                {categories.map(category => (
                  <option key={category} value={category}>
                    {category}
                  </option>
                ))}
              </select>
            </div>

            {(difficultyFilter !== 'All' || categoryFilter !== 'All') && (
              <button
                onClick={() => {
                  setDifficultyFilter('All')
                  setCategoryFilter('All')
                }}
                className="text-sm text-green-600 hover:text-green-700 font-medium transition-colors duration-200"
              >
                Clear filters
              </button>
            )}
          </div>

          {/* Results count */}
          {!loading && !error && (
            <div className="mb-5 text-sm text-gray-600 font-medium">
              Showing {filteredProblems.length} of {problems.length} problems
            </div>
          )}
          
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-5">
            {filteredProblems.map((problem) => (
              <div
                key={problem.id}
                onClick={() => handleProblemClick(problem.id)}
                className="bg-white rounded-lg p-6 cursor-pointer hover:shadow-lg transition-all duration-200 border border-gray-200 hover:border-gray-300"
                role="button"
                tabIndex={0}
                onKeyDown={(e) => {
                  if (e.key === 'Enter' || e.key === ' ') {
                    e.preventDefault()
                    handleProblemClick(problem.id)
                  }
                }}
              >
                <h2 className="text-lg font-semibold mb-4 line-clamp-2 text-gray-900 hover:text-green-600 transition-colors">{problem.title}</h2>
                
                <div className="flex items-center gap-4 mb-4 flex-wrap">
                  <span className={`font-semibold text-sm ${getDifficultyColor(problem.difficulty)}`}>
                    {problem.difficulty}
                  </span>
                  <span className="text-gray-500 text-sm">{problem.category}</span>
                </div>
                
                <div className="flex items-center justify-between text-sm text-gray-600 pt-3 border-t border-gray-100">
                  <span className="font-medium">{problem.points} points</span>
                  <span>{problem.solvedCount.toLocaleString()} solved</span>
                </div>
              </div>
            ))}
          </div>

          {/* No results message */}
          {!loading && !error && filteredProblems.length === 0 && (
            <div className="text-center py-12 text-gray-600">
              No problems match the selected filters.
            </div>
          )}
        </div>
      </div>
    </ErrorBoundary>
  )
}
