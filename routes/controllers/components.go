package controllers

const (
	home = `
    <section id="home">
    <div class="name">
		<p><small>By</small> Enzo Fernandes Francescon</p>
		</div>
        <div class="head">
            <p>Welcome to my Website!</p>
            </div>
    </section>
	`
	projects = `
	<section class="main">
        <div class="github-link">
            Click <a href="https://github.com/7Sion7/my-website" target="_blank" rel="noopener">Here</a> for the Code!
        </div>
        <div class="project-highlights">
            <h2>Here are some of my Projects:</h2>
            <ul>
                <li><div class="project">
                    <p>Recipes Website (MERN: MongoDB, Express.js, React.js, Node.js):
                        <a href="https://github.com/7Sion7/your-recipes" target="_blank" rel="noopener">github.com/7Sion7/your-recipes</a>
                    </p>

                </div></li>
                <li><div class="project">
                    <p>Forum (Go, SQLite, JavaScript, Docker, HTML and CSS):
                        <a href="https://github.com/7Sion7/forum" target="_blank" rel="noopener">github.com/7Sion7/forum</a>
                    </p>

                </div></li>
                <li><div class="project">
                    <p>Currency Converter (JavaScript, HTML and CSS):
                        --Code: <a href="https://github.com/7Sion7/currency-converter" target="_blank" rel="noopener">github.com/7Sion7/currency-converter</a>
                        --Deployed: <a href="https://7sion7.github.io/currency-converter/" target="_blank" rel="noopener">Website</a>
                    </p>

                </div></li>
                <li><div class="project">
                    <p>Artists Tracker (Go, HTML and CSS):
                        <a href="https://github.com/7Sion7/artists-tracker" target="_blank" rel="noopener">github.com/7Sion7/artists-tracker</a>
                    </p>

                </div></li>
            </ul>
        </div> 
    </section>`

	about = `
    <section class="main">
        <div class="about">
        <div class="vid-links">
        <a href="#1">About Me</a>
        <a href="#2">Professional Journey</a>
        <a href="#3">Software Development Expertise</a>
        <a href="#4">AWS Cloud and Technical Skills</a>
        <a href="#5">Professional Details</a>
        <a href="#6">Ongoing Commitment</a>
        <a href="#7">Future Aspirations</a>
        </div>
    <div class="reels">
        <div class="reel" id="1">
            <video class="videos" controls>
                <source src="static/videos/portrait-video.mp4" type="video/mp4">
                Your browser does not support the video tag.
            </video>
        </div>
        <div class="reel" id="2">
            <video class="videos" controls>
                <source src="static/videos/portrait-video.mp4" type="video/mp4">
                Your browser does not support the video tag.
            </video>
        </div>
        <div class="reel" id="3">
            <video class="videos" controls>
                <source src="static/videos/portrait-video.mp4" type="video/mp4">
                Your browser does not support the video tag.
            </video>
        </div>
        <div class="reel" id="4">
            <video class="videos" controls>
                <source src="static/videos/portrait-video.mp4" type="video/mp4">
                Your browser does not support the video tag.
            </video>
        </div>
        <div class="reel" id="5">
            <video class="videos" controls>
                <source src="static/videos/portrait-video.mp4" type="video/mp4">
                Your browser does not support the video tag.
            </video>
        </div>
        <div class="reel" id="6">
            <video class="videos" controls>
                <source src="static/videos/portrait-video.mp4" type="video/mp4">
                Your browser does not support the video tag.
            </video>
        </div>
        <div class="reel" id="7">
            <video class="videos" controls>
                <source src="static/videos/portrait-video.mp4" type="video/mp4">
                Your browser does not support the video tag.
            </video>
        </div>
    </div>
            </div>
        </section>`

	recipes = `
    <section class="project">
        <h1 class="title">Recipes</h1>
        <p>Recipes Website (MERN: MongoDB, Express.js, React.js, Node.js):
            <a href="https://github.com/7Sion7/your-recipes" target="_blank" rel="noopener">github.com/7Sion7/your-recipes</a>
        </p>
        <div class="description">
            <h3 class="tech-stack">MERN Tech Stack - (Mongo, Express.js, React.js, Node.js)</h3>
            <p>
            A Recipes WebSite has a backend implementation for a recipe management system. It provides APIs for user authentication, recipe retrieval, creation, saving, and removal. 
            The backend is built using Node.js, Express, and MongoDB. 
            Its backend is built with React. The application allows users to explore recipes, view details, and manage their saved recipes. 
            This Web Application uses the MERN stack (MongoDB, Express.js, React.js, Node.js). 
            </p>
        </div>
        <div class="images">

        </div>
    </section>`

	forum = `
    <section class="project">
        <h1 class="title">Forum</h1>
        <p>Forum (Go, SQLite, JavaScript, Docker, HTML and CSS):
            <a href="https://github.com/7Sion7/forum" target="_blank" rel="noopener">github.com/7Sion7/forum</a>
        </p>
            <div class="description">
                <h3 class="tech-stack">Tech Stack - (Go, SQLite, JavaScript, Docker)</h3>
                <p>
                This project is a forum, users can sign up, create and like posts, as well as comment on them and like the comments, the posts can be text, or images or GIFs.

                Development:
                
                -Go is used for handling requests to the server, by adding the inputted data from the client-side to the SQLite database, and sending back the data needed to the front-end;
                
                -JavaScript is used here mainly for the UI of the Web App through DOM, however, it also acts as a middleware between the client-side and the server (Go); -HTML is used for UI and making requests to the server, and displaying responses, and CSS for designing the UI aiming to make it as user-friendly as possible.
                </p>
            </div>
    <div class="images">

    </div>
    </section>
    `

	currencyConverter = `
    <section class="project">
        <h1 class="title">Currency Converter</h1>
        <p>Currency Converter (JavaScript, HTML and CSS):
            --Code: <a href="https://github.com/7Sion7/currency-converter" target="_blank" rel="noopener">github.com/7Sion7/currency-converter</a>
            --Deployed: <a href="https://7sion7.github.io/currency-converter/" target="_blank" rel="noopener">Website</a>
        </p>
        <div class="description">
            <h3 class="tech-stack">Tech Stack - (JavaScript)</h3>
            <p>
            This project is a web app currency converter written in JavaScript, this lightweight and fast web app fetches from an API the value of each 
            currency and converts the currency chosen by the user to another, also chosen by the user.
            </p>
        </div>
        <div class="images">

        </div>
    </section>
    `

	artistTracker = `
    <section class="project">
    <h1 class="title">Artist Tracker</h1>
    <p>Artists Tracker (Go, HTML and CSS):
        <a href="https://github.com/7Sion7/artists-tracker" target="_blank" rel="noopener">github.com/7Sion7/artists-tracker</a>
    </p>
    <div class="description">
        <h3 class="tech-stack">Tech Stack - (Go)</h3>
        <p>
        This program fetches data about a couple of musical artists from an API, each artist has their own page where it displays the band members, 
        its the creation date, the locations and dates of their next concerts, and their first album.
        </p>
    </div>
    <div class="images">

    </div>
</section>`
)
