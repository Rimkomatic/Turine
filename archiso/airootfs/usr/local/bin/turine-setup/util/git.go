package util

func GitClone(repo string) error {
	return Run("git", "clone", repo)
}
