document.getElementById("customersTable").addEventListener("click", function (event) {
    const target = event.target;

    if (target.tagName === "IMG" && target.getAttribute("src") === "../assets/img/icons/eye.svg") {
        const rowIndex = target.closest("tr").rowIndex;
        const table = document.getElementById("customersTable");
        const cells = table.rows[rowIndex].cells;

        // Extracting data from cells
        const id = cells[1].textContent;
        window.location.href = `/customer/details/${id}`; // Replace with your desired URL

    }
});
