const departmentsByFaculty = {
    '地域教育文化学部': ['地域教育文化学科'],
    '人文社会科学部': ['人文社会科学科'],
    '理学部': ['理学科'],
    '工学部': ['高分子・有機材料工学科', '化学・バイオ工学科', '情報・エレクトロニクス学科','機械システム工学科','建築・デザイン学科','システム創成工学科'],
    '農学部': ['食料生命環境学科'],
    '医学部': ['医学科', '看護学科']
};

function updateDepartments() {
    const facultySelect = document.getElementById('faculty');
    const departmentSelect = document.getElementById('department');
    const selectedFaculty = facultySelect.value;
    const departments = departmentsByFaculty[selectedFaculty];

    departmentSelect.innerHTML = '';

    departments.forEach(department => {
        const option = document.createElement('option');
        option.value = department;
        option.textContent = department;
        departmentSelect.appendChild(option);
    });
}

window.onload = function() {
    updateDepartments(); 
};