$(document).ready(function () {

    const dt = new Date();
    const date = dt.getMonth() + 1 + "/" + dt.getDate() + "/" + dt.getFullYear();
    $("#dateInput").val(date);

    $("#frm").on("submit", function (e) {
        e.preventDefault();
        alert("ok there");
    });

    $("#frm").on("reset", function (e) {
        setTimeout(function () { $("#dateInput").val(date); });
    });

    $("input[type=radio]").change(function () {
        const opt = this.value;

        switch(opt){
            case "modify":
                $("#submitBtn").text("Modify");
                break;
            case "search":
                $("#submitBtn").text("Search");
                break;
            default:
                $("#submitBtn").text(" Add ");
        }

    });



});