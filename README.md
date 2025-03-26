# effective_mobile_test


Env structure:
```
SERVER_ADDRESS=

POSTGRES_USER=
POSTGRES_PASSWORD=
POSTGRES_HOST=
POSTGRES_PORT=
POSTGRES_DB=
POSTGRES_CONN=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}

GOOSE_DRIVER=
GOOSE_DBSTRING=${POSTGRES_CONN}

INFO_SERVER_HOST=
```

Mock /info server is at https://7cafb837-698b-43c9-b8ac-4e6cfe8b851e.mock.pstmn.io


- [ ] Get all songs (filtration left)
- [x] Get text
- [ ] Get song
- [x] Delete song
- [x] Create song
- [x] Edit song
