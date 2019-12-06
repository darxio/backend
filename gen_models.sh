cp  internal/models/models.go ~/go/src/backend/internal/models/
cd  ~/go/src/backend/internal/models
# easyjson ~/go/src/backend/internal/models/models.go
easyjson .
cd -
cp -r ~/go/src/backend/internal/models/models_easyjson.go internal/models/