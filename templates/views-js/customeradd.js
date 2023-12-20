document.addEventListener('DOMContentLoaded', function () {
    document.getElementById("submit").addEventListener("click", function () {
        console.log("---Add Customer Button--------")
        const formData = {
            Shop : document.getElementById("shopName").value,
            Owner : document.getElementById("ownerName").value,
            PhoneNumber: document.getElementById("phoneNumber").value,
            Address: document.getElementById("address").value,
            TotalDue:  parseInt(document.getElementById("totalDue").value)  ,
        };

        console.log(formData)

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
                if (data.AddSuccess) {
                    Swal.fire({
                        icon: 'success',
                        title: 'Customer Added Successfully',
                        text: 'The Customer has been added to the database.',
                    }).then(() => {
                        window.location.href = '/customer/add';
                    });
                } else {
                    Swal.fire({
                        icon: 'error',
                        title: 'Customer Not Added',
                        text: 'There was an error adding the customer. Please try again.',
                    });
                }
            })
            .catch(error => {
                console.error("Error:", error);
            });
    })
});





