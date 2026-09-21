package cache

import "github.com/Bloem-Studios/bloem-community-theramindex-xtream-library/internal/model"

func SnapshotFromCatalog(catalog model.CatalogState) Snapshot {
	return Snapshot{Catalog: catalog, Health: catalog.Health}
}
