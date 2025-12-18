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

	if req.UserID == "" {
		res = append(res, "no UserID")
	}
	if req.Title == "" {
		res = append(res, "no Title")
	}
	if req.Description == "" {
		res = append(res, "no Description")
	}
	return res
}
