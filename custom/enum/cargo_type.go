package enum

type CargoType string

const (
	CargoTypeCargo           CargoType = "Cargo"
	CargoTypeDocuments       CargoType = "Documents"
	CargoTypeTiresWheels     CargoType = "TiresWheels"
	CargoTypePallet          CargoType = "Pallet"
	CargoTypeParcel          CargoType = "Parcel"
	CargoTypeMoney           CargoType = "Money"
	CargoTypeSignedDocuments CargoType = "SignedDocuments"
	CargoTypeTrays           CargoType = "Trays"
)
