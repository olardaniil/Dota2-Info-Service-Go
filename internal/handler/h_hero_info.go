package handler

import (
	"dota2_info_service/internal/entity"
	"net/http"
)

func (h *Handler) HeroInfoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.GetHeroInfo(w, r)
	}
}

// GetHeroInfo
// @Summary Получить информацию по заданному герою
// @Tags heroes
// @Accept json
// @Produce json
// @Param hero path string true "Имя героя"
// @Success 200 {object} entity.HeroInfoResponse
// @Failure 400 {object} Error
// @Failure 404 {object} Error
// @Router /info/{hero} [get]
func (h *Handler) GetHeroInfo(w http.ResponseWriter, r *http.Request) {
	// Получаем имя героя
	var hero = new(entity.Hero)
	hero.Name = r.URL.Path[len("/info/"):]
	// Валидация
	err := hero.Validate()
	if err != nil {
		h.response.SendError(w, 400, err.Error())
		return
	}
	// Получаем информацию о герое
	heroInfo, err := h.service.Hero.GetHeroByName(hero.Name)
	if err != nil {
		if err.Error() == "Not Found" {
			h.response.SendError(w, 404, "Герой не найден")
			return
		}
		h.response.SendError(w, 400, err.Error())
		return
	}
	hero.PopularLines = heroInfo.PopularLines
	// Отправляем ответ
	h.response.SendJSON(w, 200, hero)
}
