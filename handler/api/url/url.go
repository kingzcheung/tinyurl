package url

import (
	"encoding/json"
	"errors"
	"github.com/kingzcheung/tinyurl/config"
	"github.com/kingzcheung/tinyurl/core"
	"github.com/kingzcheung/tinyurl/handler/render"
	url2 "github.com/kingzcheung/tinyurl/store/url"
	"github.com/speps/go-hashids"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

//Shorten 短地址生成
func Shorten(urlStore core.UrlStore, hashid *hashids.HashID) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var url core.Url

		err := parseJson(r.Body, &url)
		if err != nil {
			render.Error(w, err.Error(), 400)
			return
		}

		u, err := urlStore.Create(r.Context(), &url)

		if err != nil {
			if errors.Is(err, url2.ErrDataDuplication) {
				render.Error(w, err.Error(), 400)
				return
			}
			render.Error(w, "server error", 500)
			return
		}

		if u.Slug == "" {
			slug, err := hashid.Encode([]int{u.UrlID})
			if err != nil {
				render.Error(w, err.Error(), 400)
				return
			}

			err = urlStore.UpdateSlug(r.Context(), u.UrlID, slug)
			if err != nil {
				render.Error(w, err.Error(), 400)
				return
			}
			url.Slug = slug
		}

		url.UrlID = 0
		Json(w, url)
	}
}

//List 返回地址列表
func List(urlStore core.UrlStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		query := r.URL.Query()
		searchVal := query.Get("search")

		var page, limit int
		var err error
		pageStr := query.Get("page")
		limitStr := query.Get("limit")
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			page = 1
		}

		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			limit = 10
		}

		urls, count := urlStore.List(r.Context(), core.UrlOption{
			SearchVal: searchVal,
			Page:      page,
			Limit:     limit,
		})

		w.Header().Set("X-Page-Total", strconv.FormatInt(count, 10))
		data, _ := json.Marshal(urls)

		render.Json(w, data, 200)
	}
}

func Analytics(conf *config.Config, db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//id := chi.URLParam(r, "url_id")

		//urlService := u.NewService(conf, db)
		//urlData := urlService.Find(id)
		//if urlData.UrlID == 0 {
		//	http.Error(w, "param error", http.StatusBadRequest)
		//	return
		//}
		//typ := r.URL.Query().Get("type")
		//service := tracking.NewService(r, db, urlData.UrlID)
		//data := service.ListAnalytics(typ, urlData.UrlID, 0, 0)
		//Json(w, data)
	}
}

//parseJson 从body中解析JSON
func parseJson(body io.ReadCloser, i interface{}) error {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, i)
}

func ErrJson(w http.ResponseWriter, v interface{}, status int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(v)
	_, _ = w.Write(data)
}

func Json(w http.ResponseWriter, v interface{}) {
	ErrJson(w, v, http.StatusOK)
}
