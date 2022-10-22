package comment

type Service interface {
	SaveComment(ID int, input CommentInput) (Comment, error)
	FindAllComment() ([]Comment, error)
	UpdateComment(ID int, input UpdateCommentInput) (Comment, error)
	DeleteComment(ID int) (Comment, error)
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

func (s *service) FindAllComment() ([]Comment, error) {
	comments, err := s.repository.FindAll()
	if err != nil {
		return comments, err
	}

	return comments, nil
}

func (s *service) UpdateComment(ID int, input UpdateCommentInput) (Comment, error) {
	comment, err := s.repository.FindByID(ID)
	if err != nil {
		return comment, err
	}

	comment.Message = input.Message

	updatedComment, err := s.repository.Update(comment)
	if err != nil {
		return updatedComment, err
	}

	return updatedComment, nil
}

func (s *service) DeleteComment(ID int) (Comment, error) {
	comment, err := s.repository.FindByID(ID)
	if err != nil {
		return comment, err
	}

	deletedComment, err := s.repository.Delete(comment)
	if err != nil {
		return deletedComment, err
	}

	return deletedComment, nil
}
