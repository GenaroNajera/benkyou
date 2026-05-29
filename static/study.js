let data;
let progress = document.querySelector("#progress");
let kanji = document.querySelector("#kanji");
let info = document.querySelector("#info");
let meaning = document.querySelector("#meaning");
let onyomi = document.querySelector("#onyomi");
let kunyomi = document.querySelector("#kunyomi");
let prevBtn = document.querySelector("#prevBtn");
let revealBtn = document.querySelector("#revealBtn");
let nextBtn = document.querySelector("#nextBtn");
let i = 0;

prevBtn.addEventListener("click", prev);
revealBtn.addEventListener("click", reveal);
nextBtn.addEventListener("click", next);
prevBtn.disabled = true;

// RETRIEVE FROM DATABASE
(() => {
  fetch("/api")
  .then(res => res.json())
  .then(d => {
    data = d;
    update(data[0]);
  });
})();

function update(d) {
  if (info.style.visibility !== "hidden") {
    info.style.visibility = "hidden";
    revealBtn.disabled = false;
  }

  progress.textContent = `${i + 1}/${data.length}`;
  kanji.href = `https://en.wiktionary.org/wiki/${d.Kanji}#Japanese`;
  kanji.textContent = d.Kanji;
  meaning.textContent = d.Meaning;
  onyomi.textContent = d.Onyomi;
  kunyomi.textContent = d.Kunyomi;
}

function next() {  
  i++;
  if (i == data.length - 1) {
    nextBtn.disabled = true;
  }

  if (i > 0) {
    prevBtn.disabled = false;
  }

  update(data[i]);
}

function prev() {
  i--;
  if (i == 0) {
    prevBtn.disabled = true;
  }

  if (i < data.length - 1) {
    nextBtn.disabled = false;
  }

  update(data[i]);
}

function reveal() {
  info.style.visibility = "visible";
  revealBtn.disabled = true;
}