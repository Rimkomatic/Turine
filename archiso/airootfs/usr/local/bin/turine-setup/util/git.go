package util

func GitClone(repo string) error {
	return Run("git", "clone", repo)
}



func GitCloneOn(repo string, dir string) error {
	return Run("git", "clone", repo, dir)
}
