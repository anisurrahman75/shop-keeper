document.addEventListener('DOMContentLoaded', function () {
    const viewCustomerIcons = document.querySelectorAll('.view-customer');
    viewCustomerIcons.forEach(icon => {
        icon.addEventListener('click', function (event) {
            event.preventDefault(); // Prevent the default behavior of the anchor tag
            handleView(event);
        });
    });
    function handleView(event) {
        const row = event.target.closest('tr');
        const shop = row.querySelector('#Shop').textContent.trim();
        const owner = row.querySelector('#Owner').textContent.trim();

        window.location.href = `/customer/details/${shop}/${owner}`; // Replace with your desired URL
    }
});