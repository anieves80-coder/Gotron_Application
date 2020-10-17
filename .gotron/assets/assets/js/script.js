$(document).ready(function () {
    let ws = new WebSocket("ws://localhost:" + global.backendPort + "/web/app/events");
    const dt = new Date();
    const date = dt.getMonth() + 1 + "/" + dt.getDate() + "/" + dt.getFullYear();
    $("#dateInput").val(date);

    $("#frm").on("submit", function (e) {
        e.preventDefault();
        ws.send(JSON.stringify({
            "event": "get-all"
        }));
    });

    $("#searchBtn").on("click", function (e) {
        e.preventDefault();
        alert("ok there");
    });

    $("#modifyBtn").on("click", function (e) {
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
                setModify();
                break;
            case "search":
                setSearch();
                break;
            default:
                setAdd();
        }

    });

    function setSearch() {
        $("#addBtn, #modifyBtn").addClass("btnHide");        
        $("#searchBtn").removeClass("btnHide");
        $("#dateInput, #sn2Input").val("");        
        $("#sn2Input").attr("disabled", "disabled"); 
    }
    function setAdd() {
        $("#searchBtn, #modifyBtn").addClass("btnHide");        
        $("#addBtn").removeClass("btnHide");
        $("#dateInput").val(date);
        $("#sn2Input").removeAttr("disabled");
    }
    function setModify() {
        $("#addBtn, #searchBtn").addClass("btnHide");        
        $("#modifyBtn").removeClass("btnHide");        
        $("#sn2Input").removeAttr("disabled");
        $("#dateInput").val("");
    }

    ws.onmessage = (message) => {
        let obj = JSON.parse(message.data);
        $("#tableResults").empty();
        if(obj.event === "get-all"){
            obj.eventData.forEach(element => {
                const e = JSON.parse(element);
                $("#tableResults").append(`
                    <tr>
                        <th scope="row">1</th>
                        <td>${e.rma}</td>
                        <td>${e.sn1}</td>
                        <td>${e.sn2}</td>
                        <td>${e.comment}</td>
                    </tr>                
                `);
            });
        }
    }

});
