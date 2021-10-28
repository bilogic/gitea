	"io/ioutil"
			_, err := io.Copy(ioutil.Discard, reader)
				count, err := models.Count(m)
func GetDiffRangeWithWhitespaceBehavior(gitRepo *git.Repository, beforeCommitID, afterCommitID string, maxLines, maxLineCharacters, maxFiles int, whitespaceBehavior string) (*Diff, error) {
	var cmd *exec.Cmd
		diffArgs := []string{"diff", "--src-prefix=\\a/", "--dst-prefix=\\b/", "-M"}
		cmd = exec.CommandContext(ctx, git.GitExecutable, diffArgs...)
		diffArgs := []string{"diff", "--src-prefix=\\a/", "--dst-prefix=\\b/", "-M"}
		cmd = exec.CommandContext(ctx, git.GitExecutable, diffArgs...)
	shortstatArgs := []string{beforeCommitID + "..." + afterCommitID}
func GetDiffCommitWithWhitespaceBehavior(gitRepo *git.Repository, commitID string, maxLines, maxLineCharacters, maxFiles int, whitespaceBehavior string) (*Diff, error) {
	return GetDiffRangeWithWhitespaceBehavior(gitRepo, "", commitID, maxLines, maxLineCharacters, maxFiles, whitespaceBehavior)