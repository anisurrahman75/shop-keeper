document.addEventListener('DOMContentLoaded', function () {
    document.getElementById("submit").addEventListener("click", function () {
        console.log("---Add Customer Button--------")
        const formData = {
            ShopName : document.getElementById("shopName").value,
            OwnerName : document.getElementById("ownerName").value,
            PhoneNumber: document.getElementById("phoneNumber").value,
            Address: document.getElementById("address").value,
            TotalDue: document.getElementById("totalDue").value,
        };

// Send the data to the Go backend using a POST request
        fetch('/customer/add', {
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


    // $(document).ready(function() {
    //     document.getElementById("click_modal_close").addEventListener("click", function() {
    //         const modal = document.getElementById("successModal");
    //         $(modal).modal("hide");
    //         window.location.href = "/productadd";
    //
    //     });
    // });

});





