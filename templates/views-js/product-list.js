document.addEventListener('DOMContentLoaded', function () {
    const apiUrl = '/product/list';

    const deleteProductIcons = document.querySelectorAll('.delete-product');
    deleteProductIcons.forEach(icon => {
        icon.addEventListener('click', function (event) {
            handleDelete('product', event);
        });
    });

    function handleDelete(type, event) {
        // Confirm with the user before proceeding with deletion
        Swal.fire({
            title: 'Are you sure?',
            text: 'You won\'t be able to revert this!',
            icon: 'warning',
            showCancelButton: true,
            confirmButtonColor: '#3085d6',
            cancelButtonColor: '#d33',
            confirmButtonText: 'Yes, delete it!'
        }).then((result) => {
            if (result.isConfirmed) {
                let formData={}
                const row = event.target.closest('tr');
                // Get the product name from the second column (id="productName")
                const productName = row.querySelector('#productName').textContent.trim();

                formData = {
                    Type: "product",
                    Name: productName,
                }

                console.log(formData)

                // Use fetch to send a DELETE request to your backend
                fetch(apiUrl, {
                    method: 'DELETE',
                    body: JSON.stringify(formData),
                    headers: {
                        'Content-Type': 'application/json'
                    }
                })
                    .then(response => response.json())
                    .then(data => {
                        if (data.DeleteSuccessful) {
                            // Show success alert
                            Swal.fire({
                                icon: 'success',
                                title: `${type.charAt(0).toUpperCase() + type.slice(1)} Deleted`,
                                text: `The ${type} has been deleted successfully.`,
                            }).then(() => {
                                // Reload the page after deletion
                                location.reload();
                            });
                        } else {
                            // Show error alert if deletion fails
                            Swal.fire({
                                icon: 'error',
                                title: 'Error',
                                text: `Failed to delete the ${type}. Please try again.`,
                            });
                        }
                    })
                    .catch(error => {
                        console.error('Error:', error);
                        // Show error alert if there's an issue with the fetch request
                        Swal.fire({
                            icon: 'error',
                            title: 'Error',
                            text: 'An unexpected error occurred. Please try again.',
                        });
                    });
            }
        });
    }
});
