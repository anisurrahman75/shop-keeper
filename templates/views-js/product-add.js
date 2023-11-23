document.addEventListener('DOMContentLoaded', function () {
    document.getElementById("submit").addEventListener("click", function () {
        const formData = {
            Brand : document.getElementById("brand").value,
            ProductName : document.getElementById("product_name").value,
            ProductGrade: document.getElementById("product_grade").value,
            Unit: document.getElementById("unit").value,
            Description: document.getElementById("description").value,
        };

        // Send the data to the Go backend using a POST request
        fetch('/productadd', {
            method: 'POST',
            body: JSON.stringify(formData),
            headers: {
                'Content-Type': 'application/json'
            }
        })
            .then(response => response.json())
            .then(data => {
                console.log("Response from server:", data);
                if (data.is_poduct_add_successfully){
                    console.log("Show Modal")
                    $('#successModal').modal('show');
                }
            })
            .catch(error => {
                console.error("Error:", error);
            });
    })


    $(document).ready(function() {
        document.getElementById("click_modal_close").addEventListener("click", function() {
            const modal = document.getElementById("successModal");
            $(modal).modal("hide");
            window.location.href = "/productadd";

        });
    });

});
