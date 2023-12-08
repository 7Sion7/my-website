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
        <div class="about-header">
            <h2>About Me</h2>
            <h3 class="my-name">Enzo Fernandes Francescon</h3>
            <p class="contact">Contact: <a href="https://www.linkedin.com/in/enzo-fernandes-francescon/" target="_blank" rel="noopener">LinkedIn</a> | +44 7848 817415 | enzofrancescon@gmail.com | <a href="https://github.com/7Sion7" target="_blank" rel="noopener">GitHub</a></p>
            </div>
        <p class="about-me">
        Skills
        Software Full-Stack Development (1 year experience):

        Teamwork, JavaScript (React.js, Express.js), Go, Mongo, SQLite, Docker, HTML, CSS, SCSS.
        Teaching (2 years experience):

        Languages:

        Portuguese, Spanish.
        Summary
        I am a 19-year-old student at 01 Founders, deeply passionate about developing scalable web applications through Full Stack Software Development. My goal is to join a team of seasoned developers to broaden my skills, learn from experienced software engineers, and contribute to the creation of high-quality web and mobile applications. While aspiring to be well-versed in all facets of software development, my specific interest lies in mastering AWS Cloud Engineering.

        Experience
        01 Founders
        Role:
        Planning, designing, and programming projects collaboratively, including a forum. As a team leader, I presented projects to colleagues, answered questions, and explained the development process.
        Personal Projects (Business Ideas)
        Personal Library (Tech-Stack: MERN) - Ongoing
        Utilizing AI to manage and organize user-added books by topics, facilitating syntopical reading. The application uses ChatGPT to find underlying forms in each book and create visual representations of discussions between authors. Additionally, it generates pedagogical, personalized book packages for focused study.
        Stocks Manager (Tech-Stack: MERN) - Ongoing
        Intrinsic Value calculator using Yahoo Finance API data to assist users in making informed decisions about stock shares. Users can search companies, filter financial information, view graphs, and understand relevant data for value investing.
        Other Work Experience
        Primary Music Teacher:
        Breaking down music and its aspects for Primary School children of various ages.
        Education
        GCSE:
        Maths: '7', Double Science: '7,7', English Language: '6'.
        Projects
        your-recipes
        A Recipes Website enabling users to add and save recipes, as well as fetch and save recipes from an API. Tech-Stack: MERN (MongoDB, Express.js, React.js, Node.js).
        forum
        A forum allowing users to sign up, create and like posts, categorize posts, comment on posts, and like comments. Supports text, images, and GIFs. Tech-Stack: Go, SQLite, HTML, CSS, JavaScript, Docker.
        Feel free to explore my portfolio for a deeper dive into my projects and expertise. I am enthusiastic about contributing to innovative projects and eager to continue my journey in software development.
            </p>
        </div>
        </section>`

	recipes = `
    <section class="project">
        <h1 class="title">Recipes</h1>
        <div class="description">
            <h3 class="tech-stack">MERN Tech Stack - (Mongo, Express.js, React.js, Node.js)</h3>
            <p>
            A Recipes WebSite that allows you to add your own recipes, as well as saved them and other recipes that are fetched from an API to your saved list in your account. 
            </p>
        </div>
        <div class="images">

        </div>
    </section>`

	forum = `
    <section class="project">
        <h1 class="title">Forum</h1>
            <div class="description">
                <h3 class="tech-stack">MERN Tech Stack - (Go, SQLite, JavaScript, Docker)</h3>
                <p>
                This project is a forum, users can sign up, create and like posts, as well as comment on them and like the comments, the posts can be text, or images or GIFs.
                Posts can be added to categories, and can be filtered through them, users can also see their liked posts.
                </p>
            </div>
    <div class="images">

    </div>
    </section>
    `

	currencyConverter = `
    <section class="project">
        <h1 class="title">Currency Converter</h1>
        <div class="description">
            <h3 class="tech-stack">MERN Tech Stack - (JavaScript)</h3>
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
    <div class="description">
        <h3 class="tech-stack">MERN Tech Stack - (Go)</h3>
        <p>
        This program fetches data about a couple of musical artists from an API, each artist has their own page where it displays the band members, 
        its the creation date, the locations and dates of their next concerts, and their first album.
        </p>
    </div>
    <div class="images">

    </div>
</section>`
)
