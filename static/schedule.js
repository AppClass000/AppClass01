const timeslots = document.querySelectorAll(".time-slot")

timeslots.forEach(td => {
    const classname = td.getAttribute("name");
    if (classlist[classname]) {
        td.textContent = classlist[classnane];
    };

});