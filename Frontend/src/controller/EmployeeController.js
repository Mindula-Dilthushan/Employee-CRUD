$("#btnEmpSave").click(function (event){
    let nic = $("#nic").val();
    let fullname = $("#fullname").val();
    let address = $("#address").val();
    let mobile = $("#mobile").val();

    var form = $('form').get(0);
    var data = new FormData(form);

    data.append("nic",nic);
    data.append("fullname",fullname);
    data.append("address",address);
    data.append("mobile",mobile);

    $.ajax(
        {
            type: "POST",
            enctype: 'multipart/form-data',
            url: "http://localhost:8001/api/employee",
            data: data,
            processData: false,
            contentType: false,
            cache: false,
            success: function (data) {
                alert("Success");
            },
            error: function (e) {
                alert("Error Please Try again");
            }
        }
    );
});