# english flashcards

## api endpoints
- `POST` `/v1/auth/login`

Request:
```json
{
    "email": "test@test.com",
    "password": "test"
}
```

Response:

`200 (OK)`
```json
{
	"userId": 1,
	"email": "test@test.com",
	"password": "test",
	"session": "JJbTuZkkskZpLIfKk2ds",
	"expires_at": "2022-11-21T02:55:19.134387Z"
}
```

- `POST` `/v1/auth/register`

Request:

```json
{
    "email": "test@test.com",
    "password": "test"
}
```
Response:

`201 (Created)`

- `GET` `/v1/decks/{deckID}`

Response:

`200 (OK)`

```json
[
	{
		"id": 9,
		"deck_id": 1,
		"english": "test",
		"russian": "test",
		"association": "test",
		"example": null,
		"transcription": null,
		"image": null,
		"image_url": null,
		"audio": null,
		"audio_url": null,
		"created_at": "2022-10-22T01:51:01.888931Z"
	},
	{
		"id": 10,
		"deck_id": 1,
		"english": "test",
		"russian": "test",
		"association": null,
		"example": null,
		"transcription": null,
		"image": null,
		"image_url": null,
		"audio": null,
		"audio_url": null,
		"created_at": "2022-10-22T01:52:54.688469Z"
	}
]
```

- `POST` `/v1/cards`

Request:

```multipart
--boundary
Content-Disposition: form-data; name="english"
test
--boundary
Content-Disposition: form-data; name="russian"
test
--boundary
Content-Disposition: form-data; name="example"
test
--boundary
Content-Disposition: form-data; name="image"; filename="example.jpg"
test
--boundary--
```

Response:

`201 (Created)`

```json
{
	"ID": 13,
	"DeckID": 1,
	"English": "test",
	"Russian": "test",
	"Association": null,
	"Example": null,
	"Transcription": null,
	"Image": null,
	"ImageUrl": null,
	"Audio": null,
	"AudioUrl": null,
	"created_at": "2022-10-22T02:01:40.886515Z"
}
```
