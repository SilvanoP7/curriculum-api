package handlers

import (
	"net/http"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/SilvanoP7/curriculum-api/models"
	_ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "postgres"
    dbname   = "curriculum"
)

// Pong             godoc
// @Summary      Basic health check to ensure service is responding
// @Description  Responds with pong status.
// @Tags         health
// @Produce      json
// @Success      200 {object}   models.Pong
// @Router       /ping [get]
func Pong(c *gin.Context) {
        var pong models.Pong
        pong.Status = "pong"
	c.JSON(http.StatusOK, &pong)
}

// Pong             godoc
// @Summary      Get all subjects
// @Description  Responds with all subjects
// @Tags         getSubjects
// @Produce      json
// @Success      200 {array}   models.Subjects
// @Router       /getSubjects [get]
func GetSubjects(c *gin.Context) {
        var fail models.Pong
        psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	rowsRs, err := db.Query("SELECT subject_id, subject, key_stage, purpose_of_study FROM subjects")

	if err != nil {
                fail.Status = "Failed getting data"
		c.JSON(http.StatusInternalServerError, &fail)
		return
	}
	defer rowsRs.Close()

	snbs := make([]models.Subjects, 0)


	// we loop through the values of rows
	for rowsRs.Next() {
		snb := models.Subjects{}
		err := rowsRs.Scan(&snb.SubjectID, &snb.Subject, &snb.KeyStage, &snb.PurposeOfStudy)
		if err != nil {
	                fail.Status = "Failed getting data"
        	        c.JSON(http.StatusInternalServerError, &fail)
			return
		}
		snbs = append(snbs, snb)
	}

	if err = rowsRs.Err(); err != nil {
		c.JSON(http.StatusOK, &snbs)
		return
	}

	db.Close()
	c.JSON(http.StatusOK, &snbs)       
}

// Pong             godoc
// @Summary      Get Subject Contents
// @Description  Responds with the all subject contents
// @Tags         getSubjectContent
// @Produce      json
// @Success      200 {array}   models.SubjectContent
// @Router       /getSubjectContent [get]
func GetSubjectContent(c *gin.Context) {
        var fail models.Pong
        psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
        db, err := sql.Open("postgres", psqlInfo)

        rowsRs, err := db.Query("SELECT subject_id, subject_content, subject_content_id FROM subject_content")

        if err != nil {
                fail.Status = "Failed getting data"
                c.JSON(http.StatusInternalServerError, &fail)
                return
        }
        defer rowsRs.Close()

        snbs := make([]models.SubjectContent, 0)


        // we loop through the values of rows
        for rowsRs.Next() {
                snb := models.SubjectContent{}
                err := rowsRs.Scan(&snb.SubjectID, &snb.SubjectContent, &snb.SubjectContentId)
                if err != nil {
                        fail.Status = "Failed getting data"
                        c.JSON(http.StatusInternalServerError, &fail)
                        return
                }
                snbs = append(snbs, snb)
        }

        if err = rowsRs.Err(); err != nil {
                c.JSON(http.StatusOK, &snbs)
                return
        }

        db.Close()
        c.JSON(http.StatusOK, &snbs)
}


// Pong             godoc
// @Summary      Basic health check to ensure the service can connect to the db
// @Description  Responds with fail or connected status
// @Tags         DbTest
// @Produce      json
// @Success      200 {object}   models.DbTest
// @Router       /db [get]
func DbTest(c *gin.Context) {
  var dbt models.DbTest
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    dbt.Status = "Fail"
    c.JSON(http.StatusInternalServerError,&dbt)
    return
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
    dbt.Status = "Fail"
    c.JSON(http.StatusInternalServerError,&dbt)
    return
  }
  dbt.Status = "Connected"
  c.JSON(http.StatusOK,&dbt)
}
