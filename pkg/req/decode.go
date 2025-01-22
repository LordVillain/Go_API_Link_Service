package req

import (
	"encoding/json"
	"io"
)

func Decode[T any](body io.ReadCloser) (T, error) {
	var payload T
	err := json.NewDecoder(body).Decode(&payload)
	if err != nil {
		return payload, err
	}
	return payload, nil
}

    // if payload.Email == "" {
	// 	res.Json(w, "Email required", 402)
	// 	return
	// }
	// // если частая операция то
	// // reg, _ := regexp.Compile(`[a-zA-Z0-9\._%+\-]+@[a-zA-Z0-9\.\-]+\.[a-zA-Z]{2,}`)
	// // if !reg.MatchString(payload.Email){
	// // 	res.Json(w, "Wrong email", 402)
	// // 	return
	// // }
	// // еще способ:
	// // _, err = mail.ParseAddress(payload.Email)
	// // if err != nil {
	// // 	res.Json(w, "Wrong email", 402)
	// // 	return
	// // }
	// // если не часто то:
	// match, _ := regexp.MatchString(`[a-zA-Z0-9\._%+\-]+@[a-zA-Z0-9\.\-]+\.[a-zA-Z]{2,}`, payload.Email)
	// if !match {
	// 	res.Json(w, "Wrong email", 402)
	// 	return
	// }
	// if payload.Password == "" {
	// 	res.Json(w, "Password required", 402)
	// 	return
	// }
