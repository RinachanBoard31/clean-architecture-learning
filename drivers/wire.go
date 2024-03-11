// [4] Frameworks & Drivers

package drivers

// "clean-architecture-learning/adapters/controllers"
// "clean-architecture-learning/adapters/gateways"
// "clean-architecture-learning/adapters/presenters"
// "clean-architecture-learning/database"
// "clean-architecture-learning/usecases/interactors"
// "context"

// "github.com/google/wire"
// "github.com/labstack/echo/v4"

// func InitializeUserDriver(ctx context.Context) (User, error) {
// 	wire.Build(
// 		NewFirestoreClientFactory,
// 		echo.New,
// 		NewOutputFactory,
// 		NewInputFactory,
// 		NewRepositoryFactory,
// 		controllers.NewUserController,
// 		NewUserDriver,
// 	)
// 	return &UserDriver{}, nil
// }

// func NewFirestoreClientFactory() gateways.FirestoreClientFactory {
// 	return &database.MyFirestoreClientFactory{}
// }

// func NewOutputFactory() controllers.OutputFactory {
// 	return presenters.NewUserOutputPort
// }

// func NewInputFactory() controllers.InputFactory {
// 	return interactors.NewUserInputPort
// }

// func NewRepositoryFactory() controllers.RepositoryFactory {
// 	return gateways.NewUserRepository
// }
