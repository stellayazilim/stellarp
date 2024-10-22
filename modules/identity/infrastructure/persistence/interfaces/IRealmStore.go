package interfaces

import "StellaRP/modules/identity/domain/realm"

type IRealmStore interface {
	IStore[realm.IRealm]
}
