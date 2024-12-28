const dropzone = document.getElementById('dropzone');
const convertBtn = document.getElementById('convertBtn');
const formatSelect = document.getElementById('formatSelect');
const progressBar = document.getElementById('progressBar');
const progressContainer = document.getElementById('progressContainer');
const modal = document.getElementById('modal');
const downloadBtn = document.getElementById('downloadBtn');
const closeModal = document.getElementById('closeModal');
const themeSwitcher = document.getElementById('themeSwitcher');

const formatMapping = {
    pages:   ['pdf', 'docx'],
    numbers: ['pdf'],
    key:     ['pdf'],
    docx:    ['pdf'],
    pdf:     ['docx']
};

let selectedFile = null;
let downloadUrl = '';

window.addEventListener('load', () => {
    const savedTheme = localStorage.getItem('theme');
    if (savedTheme === 'dark') {
        document.body.classList.add('dark-theme');
        themeSwitcher.checked = true;
    } else {
        document.body.classList.remove('dark-theme');
        themeSwitcher.checked = false;
    }
});

themeSwitcher.addEventListener('change', () => {
    const isDark = themeSwitcher.checked;
    document.body.classList.toggle('dark-theme', isDark);
    localStorage.setItem('theme', isDark ? 'dark' : 'light');
});


dropzone.addEventListener('click', () => {
    document.getElementById('fileInput').click();
});

document.getElementById('fileInput').addEventListener('change', (e) => {
    selectedFile = e.target.files[0];
    dropzone.innerText = selectedFile.name;

    const ext = selectedFile.name.split('.').pop().toLowerCase();

    if (formatMapping[ext]) {
        updateFormatOptions(formatMapping[ext]);
    } else {
        alert('Формат файла неизвестен или не поддерживается!');
    }
});

function updateFormatOptions(options) {
    const select = document.getElementById('formatSelect');
    select.innerHTML = ''; // Очищаем варианты

    options.forEach((opt) => {
        const option = document.createElement('option');
        option.value = opt;
        option.textContent = opt.toUpperCase();
        select.appendChild(option);
    });
}

convertBtn.addEventListener('click', async () => {
    if (!selectedFile) {
        alert('Please upload a file!');
        return;
    }

    progressContainer.classList.remove('hidden');

    const formData = new FormData();
    formData.append('file', selectedFile);

    const format = formatSelect.value;
    const url = `/convert?to=${format}`;

    try {
        const response = await fetch(url, { method: 'POST', body: formData });
        if (!response.ok) {
            const errorText = await response.text();
            alert(`Conversion failed: ${errorText}`);
            progressContainer.classList.add('hidden');
            return;
        }

        const blob = await response.blob();
        downloadUrl = window.URL.createObjectURL(blob);
        downloadBtn.href = downloadUrl;
        downloadBtn.download = 'converted-file';

        progressContainer.classList.add('hidden');
        modal.classList.remove('hidden');
    } catch (err) {
        alert('Error: ' + err.message);
        progressContainer.classList.add('hidden');
    }
});

closeModal.addEventListener('click', () => {
    modal.classList.add('hidden');
});