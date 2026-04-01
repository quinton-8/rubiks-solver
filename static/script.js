// Map letters to standard Rubik's colors
const COLOR_MAP = {
    'U': '#FFFFFF', // Up = White
    'R': '#B71234', // Right = Red
    'F': '#009B48', // Front = Green
    'D': '#FFD500', // Down = Yellow
    'L': '#FF5800', // Left = Orange
    'B': '#0046AD'  // Back = Blue
};

// The exact order Kociemba needs the string
const FACE_ORDER = ['U', 'R', 'F', 'D', 'L', 'B'];
let selectedColorKey = 'U';

// Initialize the UI
function init() {
    buildPalette();
    buildGrid();
    resetCube();

    document.getElementById('btn-reset').addEventListener('click', resetCube);
    document.getElementById('btn-solve').addEventListener('click', solveCube);
}

function buildPalette() {
    const palette = document.getElementById('palette');
    for (const [key, color] of Object.entries(COLOR_MAP)) {
        const swatch = document.createElement('div');
        swatch.className = 'color-swatch';
        swatch.style.backgroundColor = color;
        swatch.dataset.key = key;
        
        if (key === selectedColorKey) swatch.classList.add('active');
        
        swatch.addEventListener('click', () => {
            document.querySelectorAll('.color-swatch').forEach(s => s.classList.remove('active'));
            swatch.classList.add('active');
            selectedColorKey = key;
        });
        palette.appendChild(swatch);
    }
}

function buildGrid() {
    FACE_ORDER.forEach(faceKey => {
        const faceDiv = document.getElementById(`face-${faceKey}`);
        for (let i = 0; i < 9; i++) {
            const facelet = document.createElement('div');
            facelet.className = 'facelet';
            facelet.dataset.face = faceKey;
            facelet.dataset.index = i;
            
            // The 5th square (index 4) is the center, it cannot be colored
            if (i === 4) {
                facelet.classList.add('center');
                facelet.dataset.colorKey = faceKey;
                facelet.style.backgroundColor = COLOR_MAP[faceKey];
            } else {
                facelet.addEventListener('click', (e) => {
                    e.target.dataset.colorKey = selectedColorKey;
                    e.target.style.backgroundColor = COLOR_MAP[selectedColorKey];
                });
            }
            faceDiv.appendChild(facelet);
        }
    });
}

function resetCube() {
    FACE_ORDER.forEach(faceKey => {
        const facelets = document.getElementById(`face-${faceKey}`).children;
        for (let i = 0; i < 9; i++) {
            facelets[i].dataset.colorKey = faceKey;
            facelets[i].style.backgroundColor = COLOR_MAP[faceKey];
        }
    });
    document.getElementById('result').textContent = "Awaiting input...";
    document.getElementById('result').className = "";
}

// Scrape the grid and send the 54-char string to the Go Server
async function solveCube() {
    const resultDiv = document.getElementById('result');
    resultDiv.textContent = "Solving in background...";
    resultDiv.className = "";

    let cubeString = "";
    // MUST be read in U, R, F, D, L, B order for Kociemba
    FACE_ORDER.forEach(faceKey => {
        const facelets = document.getElementById(`face-${faceKey}`).children;
        for (let i = 0; i < 9; i++) {
            cubeString += facelets[i].dataset.colorKey;
        }
    });

    try {
        const response = await fetch('/api/solve', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ cubeString: cubeString })
        });

        const data = await response.json();

        if (data.error) {
            resultDiv.textContent = `Error: ${data.error}`;
            resultDiv.className = "error";
        } else {
            resultDiv.textContent = data.solution;
            resultDiv.className = "success";
        }
    } catch (err) {
        resultDiv.textContent = "Server connection failed. Is the Go server running?";
        resultDiv.className = "error";
    }
}

window.onload = init;