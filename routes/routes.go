package routes

import "go.uber.org/fx"

// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(NewRoutes),
	//fx.Provide(NewUtilityRoutes),
	fx.Provide(NewUserRoutes),
	fx.Provide(NewQuizRoutes),
	fx.Provide(NewQuestionRoutes),
	fx.Provide(NewChoicesRoutes),
	fx.Provide(NewHistoryRoutes),
	fx.Provide(NewClientRoutes),
	fx.Provide(NewFriendRequestRoutes),
	fx.Provide(NewConversationRoutes),
	//fx.Provide(NewPlantRoutes),
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// NewRoutes sets up routes
func NewRoutes(
	//utilityRoutes UtilityRoutes,
	userRoutes UserRoutes,
	//plantRoutes PlantRoutes,
	quizRoutes QuizRoutes,
	questonRoutes QuestionsRoutes,
	choiceRoutes ChoiceRoutes,
	historyRoutes HistoryRoutes,
	clientRoutes ClientRoutes,

) Routes {
	return Routes{
		//utilityRoutes,
		userRoutes,
		quizRoutes,
		questonRoutes,
		choiceRoutes,
		historyRoutes,
		clientRoutes,

		//plantRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()

	}
}
