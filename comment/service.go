package comment

type Service interface {
	SaveComment(ID int, input CommentInput) (Comment, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) SaveComment(ID int, input CommentInput) (Comment, error) {
	comment := Comment{}

	comment.UserID = ID
	comment.PhotoID = input.PhotoID
	comment.Message = input.Message

	newComment, err := s.repository.Create(comment)
	if err != nil {
		return newComment, err
	}
	return newComment, nil
}
