const projectsBtn = document.getElementById('projects');
const dropdownDivProjects = document.getElementById('projects-dropdown');
const projectsDropdown = document.getElementById('dropdown');
const dropdownLinks =  projectsDropdown.querySelectorAll('a');
const vidBtns = document.querySelectorAll('.vid-btn');
const videos = document.querySelectorAll('video');
const aLinks = document.querySelectorAll('a')



aLinks.forEach((a)=> {
    a.addEventListener('click', ()=> {
        const vidDiv = document.getElementById(a.href.slice(-1));
        const selectedVid = vidDiv.querySelector('video');
        manageVideos();
        selectedVid.play();
        console.log(selectedVid)
    })
});

const manageVideos = () => {
    videos.forEach((video)=> {
        video.pause();
    })
};

const fadeIn = () => {
    projectsDropdown.style.pointerEvents = 'all';
    projectsBtn.style.color = 'white';
    projectsDropdown.style.opacity = '1';
    dropdownLinks.forEach(link => {
        if (link.classList.contains("fade-out"))link.classList.remove("fade-out")
        link.classList.replace("fades", "fade-in")
    })
}

const fadeOut = () => {
    projectsDropdown.style.pointerEvents = 'none';
    projectsBtn.style.color = 'var(--lightWhite)';
    projectsDropdown.style.opacity = '0';
    dropdownLinks.forEach(link => {
        link.classList.replace("fade-in", "fade-out")
        link.classList.add("fades")
    })
}

const removeHead = (htmlString) => {
    const parser = new DOMParser();
    const doc = parser.parseFromString(htmlString, 'text/html');

    const headElement = doc.querySelector('head');
    if (headElement) {
        headElement.remove();
    }
    const modifiedHtmlString = new XMLSerializer().serializeToString(doc);

    return modifiedHtmlString;
}

const go = async (id) => {
    await fetch(`/${id}`, {method: 'GET',})
        .then(response => response.text())
        .then(data => {
            document.body.innerHTML = removeHead(data);
        })
        .catch(error => console.error('Error fetching data:', error));
}