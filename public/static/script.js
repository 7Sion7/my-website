const projectsBtn = document.getElementById('projects');
const dropdownDivProjects = document.getElementById('projects-dropdown');
const projectsDropdown = document.getElementById('dropdown');
const dropdownLinks =  projectsDropdown.querySelectorAll('a');

// projectsBtn.addEventListener('mouseover', function () {
//     console.log("Hey")
//     fadeIn();
// });

// dropdownDivProjects.addEventListener('mouseleave', function (event) {
//     if (!event.relatedTarget || !event.relatedTarget.closest('#projects-dropdown')) {
//         console.log("Hi")
//         fadeOut();
//     }
// });

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