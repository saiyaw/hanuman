function save_candidate_info(){
    var result = null;
    $.ajax({
        type: "POST",
        async: false,
        url: "/insertcandidate",
        data: {
            "fullname": $("#iname").val(),
            "age": $("#iage").val(),
            "gender": $("#sgender").val(),
            "mobile": $("#imobile").val(),
            "email": $("#iemail").val(),
            "workyear": $("#iworkyear").val(),
        },
        success: function(r) {
            result = r;
        }
    });
    return result;
}

/*

func (c *CandidateController) InsertOneCandidate() {
    var candidate models.Candidate
    candidate.Fullname = c.GetString("fullname")

    age, _ := c.GetInt64("age")
    candidate.Age = age

    candidate.Gender = c.GetString("gender")

    candidate.Mobile = c.GetString("mobile")

    candidate.Email = c.GetString("email")

    workyear, _ := c.GetInt64("workyear")
    candidate.Workyear = workyear

    candidate.Insert()
    c.Ctx.WriteString(strconv.FormatInt(candidate.Id, 10))

}


*/