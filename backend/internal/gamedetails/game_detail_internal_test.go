package gamedetails

import (
	"database/sql"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GameDetailInternal", func() {
	var (
		db      *sql.DB
		mock    sqlmock.Sqlmock
		handler GameDetailInternalHandler
	)

	BeforeEach(func() {
		var err error
		db, mock, err = sqlmock.New()
		Expect(err).NotTo(HaveOccurred())
		handler = GameDetailInternalHandler{DB: db}
	})

	AfterEach(func() {
		db.Close()
	})

	Describe("GetGameDetailRecord", func() {
		Context("when the record exists", func() {
			It("should return the game detail record", func() {
				igdbID := 123
				expectedTime := time.Now()

				rows := sqlmock.NewRows([]string{
					"id", "igdb_id", "is_owned", "is_beaten", "is_wishlisted",
					"name", "thumbnail_url", "cover_url", "created_at", "updated_at",
				}).AddRow(1, igdbID, true, false, true, "Test Game", "thumb.jpg", "cover.jpg", expectedTime, expectedTime)

				mock.ExpectQuery("SELECT (.+) FROM game_details WHERE igdb_id = ?").
					WithArgs(igdbID).
					WillReturnRows(rows)

				result, err := handler.GetGameDetailRecord(igdbID)

				Expect(err).NotTo(HaveOccurred())
				Expect(result).NotTo(BeNil())
				Expect(result.ID).To(Equal(1))
				Expect(result.IGDBID).To(Equal(igdbID))
				Expect(result.IsOwned).To(BeTrue())
				Expect(result.IsBeaten).To(BeFalse())
				Expect(result.IsWishlisted).To(BeTrue())
				Expect(result.Name).To(Equal("Test Game"))
				Expect(result.ThumbnailURL).To(Equal("thumb.jpg"))
				Expect(result.CoverURL).To(Equal("cover.jpg"))
			})
		})

		Context("when the record does not exist", func() {
			It("should return empty record without error", func() {
				igdbID := 123

				mock.ExpectQuery("SELECT (.+) FROM game_details WHERE igdb_id = ?").
					WithArgs(igdbID).
					WillReturnError(sql.ErrNoRows)

				result, err := handler.GetGameDetailRecord(igdbID)

				Expect(err).NotTo(HaveOccurred())
				Expect(result).NotTo(BeNil())
			})
		})

		Context("when igdbID is 0", func() {
			It("should return an error", func() {
				result, err := handler.GetGameDetailRecord(0)

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("id must be provided to get details"))
				Expect(result).To(BeNil())
			})
		})
	})

	Describe("UpsertGameDetailRecord", func() {
		Context("when creating a new record", func() {
			It("should create a new game detail record", func() {
				igdbID := 123
				name := "Test Game"
				isOwned := true
				isBeaten := false
				isWishlisted := true
				thumbnailURL := "thumb.jpg"
				coverURL := "cover.jpg"

				input := UpsertGameDetailInput{
					IGDBID:       &igdbID,
					Name:         &name,
					IsOwned:      &isOwned,
					IsBeaten:     &isBeaten,
					IsWishlisted: &isWishlisted,
					ThumbnailURL: &thumbnailURL,
					CoverURL:     &coverURL,
				}

				mock.ExpectExec("INSERT INTO game_details").
					WithArgs(
						igdbID,
						isOwned,
						isBeaten,
						isWishlisted,
						name,
						thumbnailURL,
						coverURL,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))

				// Mock the subsequent GetGameDetailRecord call
				expectedTime := time.Now()
				rows := sqlmock.NewRows([]string{
					"id", "igdb_id", "is_owned", "is_beaten", "is_wishlisted",
					"name", "thumbnail_url", "cover_url", "created_at", "updated_at",
				}).AddRow(1, igdbID, isOwned, isBeaten, isWishlisted, name, thumbnailURL, coverURL, expectedTime, expectedTime)

				mock.ExpectQuery("SELECT (.+) FROM game_details WHERE igdb_id = ?").
					WithArgs(igdbID).
					WillReturnRows(rows)

				result, err := handler.UpsertGameDetailRecord(input)

				Expect(err).NotTo(HaveOccurred())
				Expect(result).NotTo(BeNil())
				Expect(result.IGDBID).To(Equal(igdbID))
				Expect(result.Name).To(Equal(name))
				Expect(result.IsOwned).To(Equal(isOwned))
			})
		})

		Context("when updating an existing record", func() {
			It("should update the existing game detail record", func() {
				id := 1
				igdbID := 123
				name := "Updated Game"
				isOwned := true
				isBeaten := true
				isWishlisted := false
				thumbnailURL := "new_thumb.jpg"
				coverURL := "new_cover.jpg"

				input := UpsertGameDetailInput{
					ID:           &id,
					IGDBID:       &igdbID,
					Name:         &name,
					IsOwned:      &isOwned,
					IsBeaten:     &isBeaten,
					IsWishlisted: &isWishlisted,
					ThumbnailURL: &thumbnailURL,
					CoverURL:     &coverURL,
				}

				// Mock the initial GetGameDetailRecord call
				expectedTime := time.Now()
				rows := sqlmock.NewRows([]string{
					"id", "igdb_id", "is_owned", "is_beaten", "is_wishlisted",
					"name", "thumbnail_url", "cover_url", "created_at", "updated_at",
				}).AddRow(id, igdbID, false, false, true, "Old Name", "old_thumb.jpg", "old_cover.jpg", expectedTime, expectedTime)

				mock.ExpectQuery("SELECT (.+) FROM game_details WHERE igdb_id = ?").
					WithArgs(igdbID).
					WillReturnRows(rows)

				mock.ExpectExec("UPDATE game_details").
					WithArgs(
						isOwned,
						isBeaten,
						isWishlisted,
						name,
						thumbnailURL,
						coverURL,
						id,
						igdbID,
					).
					WillReturnResult(sqlmock.NewResult(0, 1))

				// Mock the final GetGameDetailRecord call
				rows = sqlmock.NewRows([]string{
					"id", "igdb_id", "is_owned", "is_beaten", "is_wishlisted",
					"name", "thumbnail_url", "cover_url", "created_at", "updated_at",
				}).AddRow(id, igdbID, isOwned, isBeaten, isWishlisted, name, thumbnailURL, coverURL, expectedTime, expectedTime)

				mock.ExpectQuery("SELECT (.+) FROM game_details WHERE igdb_id = ?").
					WithArgs(igdbID).
					WillReturnRows(rows)

				result, err := handler.UpsertGameDetailRecord(input)

				Expect(err).NotTo(HaveOccurred())
				Expect(result).NotTo(BeNil())
				Expect(result.ID).To(Equal(id))
				Expect(result.IGDBID).To(Equal(igdbID))
				Expect(result.Name).To(Equal(name))
				Expect(result.IsOwned).To(Equal(isOwned))
				Expect(result.IsBeaten).To(Equal(isBeaten))
			})
		})

		Context("when IGDBID is not provided", func() {
			It("should return an error", func() {
				input := UpsertGameDetailInput{}

				result, err := handler.UpsertGameDetailRecord(input)

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("must provide IGDBID in order to save details"))
				Expect(result).To(BeNil())
			})
		})
	})
})
