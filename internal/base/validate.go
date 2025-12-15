package base

type ValidateRequest struct {
	UserID      string
	Title       string
	Description string
}

func Validate(req *ValidateRequest) []string {
	res := make([]string, 0)

	if req == nil {
		res = append(res, "no data")
		return res
	}

	userId := req.UserID
	Title := req.Title
	Description := req.Description
	if userId == "" {
		res = append(res, "no UserID")
	}
	if Title == "" {
		res = append(res, "no Title")
	}
	if Description == "" {
		res = append(res, "no Description")
	}
	return res
}
