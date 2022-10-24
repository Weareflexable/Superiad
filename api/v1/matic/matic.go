package matic

import (
	"github.com/TheLazarusNetwork/superiad/api/middleware/onlyunlockedmiddleware"
	"github.com/TheLazarusNetwork/superiad/api/v1/matic/approve"
	approveall "github.com/TheLazarusNetwork/superiad/api/v1/matic/approveAll"
	"github.com/TheLazarusNetwork/superiad/api/v1/matic/checkbalance"
	"github.com/TheLazarusNetwork/superiad/api/v1/matic/delegate"
	"github.com/TheLazarusNetwork/superiad/api/v1/matic/fetchwallet"
	"github.com/TheLazarusNetwork/superiad/api/v1/matic/getstatus"
	"github.com/TheLazarusNetwork/superiad/api/v1/matic/isowner"
	"github.com/TheLazarusNetwork/superiad/api/v1/matic/setstatus"
	signmessage "github.com/TheLazarusNetwork/superiad/api/v1/matic/signMessage"
	"github.com/TheLazarusNetwork/superiad/api/v1/matic/tokenuri"
	"github.com/TheLazarusNetwork/superiad/api/v1/matic/transfer"
	verifysignature "github.com/TheLazarusNetwork/superiad/api/v1/matic/verifySignature"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes Use the given Routes
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/matic")
	{
		checkbalance.ApplyRoutes(v1)
		fetchwallet.ApplyRoutes(v1)
		isowner.ApplyRoutes(v1)
		verifysignature.ApplyRoutes(v1)
		tokenuri.ApplyRoutes(v1)
		getstatus.ApplyRoutes(v1)
		setstatus.ApplyRoutes(v1)

		v1.Use(onlyunlockedmiddleware.OnlyUnlocked())
		signmessage.ApplyRoutes(v1)
		transfer.ApplyRoutes(v1)
		approve.ApplyRoutes(v1)
		approveall.ApplyRoutes(v1)
		delegate.ApplyRoutes(v1)
	}
}
