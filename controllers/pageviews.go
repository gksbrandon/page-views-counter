package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gksbrandon/page-views-counter/models"
	"github.com/gksbrandon/page-views-counter/utils"

	"github.com/gorilla/mux"
)

// Error checking
func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Verify within time span
func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

type Controller struct{}

// Get number of views for certain time periods
func (c Controller) GetView(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Initialize variables
		var view models.SingleView
		var PageViews []models.SingleView
		var GetViewResponse models.GetView

		// Obtain article id from URL
		params := mux.Vars(r)

		// Query postgres database for selected id within past 3 days
		sqlStatement := `SELECT * FROM pages 
			WHERE article_id = $1 AND t_timestamp BETWEEN NOW() - INTERVAL '72 HOURS' AND NOW() 
			ORDER BY t_timestamp DESC`
		rows, err := db.Query(sqlStatement, params["article_id"])
		logFatal(err)
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&view.Id, &view.ArticleID, &view.TimeStamp)
			logFatal(err)
			PageViews = append(PageViews, view)
		}

		// Initialize count variables
		var fiveMinutesAgoCount int
		var oneHourAgoCount int
		var oneDayAgoCount int
		var twoDaysAgoCount int
		var threeDaysAgoCount int

		// Increment count variables depending on number of views within time frames
		now := time.Now()
		for _, value := range PageViews {
			if inTimeSpan(now.Add(time.Duration(-5)*time.Minute), now, value.TimeStamp) {
				fiveMinutesAgoCount++
				oneHourAgoCount++
				oneDayAgoCount++
				twoDaysAgoCount++
				threeDaysAgoCount++
			} else if inTimeSpan(now.Add(time.Duration(-1)*time.Hour), now, value.TimeStamp) {
				oneHourAgoCount++
				oneDayAgoCount++
				twoDaysAgoCount++
				threeDaysAgoCount++
			} else if inTimeSpan(now.Add(time.Duration(-24)*time.Hour), now, value.TimeStamp) {
				oneDayAgoCount++
				twoDaysAgoCount++
				threeDaysAgoCount++
			} else if inTimeSpan(now.Add(time.Duration(-48)*time.Hour), now, value.TimeStamp) {
				twoDaysAgoCount++
				threeDaysAgoCount++
			} else if inTimeSpan(now.Add(time.Duration(-72)*time.Hour), now, value.TimeStamp) {
				threeDaysAgoCount++
			}
		}

		// Creating our response object
		count := []models.Count{}
		count = append(count,
			models.Count{"5 minutes ago", fiveMinutesAgoCount},
			models.Count{"1 hour ago", oneHourAgoCount},
			models.Count{"1 day ago", oneDayAgoCount},
			models.Count{"2 days ago", twoDaysAgoCount},
			models.Count{"3 days ago", threeDaysAgoCount})

		GetViewResponse.Data.ArticleID = params["article_id"]
		GetViewResponse.Data.Type = "statistics_article_view_count"
		GetViewResponse.Data.Attributes.Count = count

		// Successful response
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, GetViewResponse)
	}
}

// Add a view to the database
func (c Controller) AddView(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Initialize variables
		var view models.AddView
		err := json.NewDecoder(r.Body).Decode(&view)
		logFatal(err)

		// Query postgres database, inserting a new view
		sqlStatement := `INSERT INTO pages (article_id, t_timestamp) 
			VALUES ($1, $2) 
			RETURNING id;`
		t := time.Now()
		id := 0
		err = db.QueryRow(sqlStatement, view.Data.ArticleID, t).Scan(&id)
		logFatal(err)

		// Successful response
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, view)
	}
}
