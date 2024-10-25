"use strict";

class screenState {
  title;
  displayText;
  github;

  constructor(title, displayText, github_address) {
    this.title = title;
    this.displayText = displayText;
    this.github_address = github_address;
  }
}

const ascii = document.querySelector("#ascii-art");
const jobTitle = document.querySelector("#job-title");
const displayArea = document.querySelector('#display-area');
const contentSpace = document.createElement("p");
const title = document.createElement("h1");
// const github = document.createElement("a");

let projectsHub = false;
let browsingProjects = false; 
let projectIndex = 0;
let home = true;

const homeScreen = new screenState("", "Press P for PROJECTS | Press B for BIO | Press H to return HOME", "");
const projectHome = new screenState("Projects", "Use the right and left arrow keys to navigate the projects. Press G to view a project's Github page.", "")
const project1Screen = new screenState("The Isles", "A prototype stage first-person RPG written from scratch in C++. Real time 3D rendering using D3D11, audio using WASAPI.", "https://github.com/jpres27/learn_d3d11");
const project2Screen = new screenState("The Libra Project", "Contributed audit querying functionality designed to work with AWS Lambda to the Libra redesign project at the University of Virginia Library. Libra is a scholarly repository.", "https://github.com/uvalib/libra-lambda/tree/main/libra-audit-query");
const project3Screen = new screenState("8086 Instruction Decoder", "Decodes 8086 machine instructions and provides a valid 8086 assembly output file. The output can be fed to an assembler and produce identical machine code. This was created as part of Casey Muratori's Performance Aware Programming course and involved many hours of nostalgic consultation of Intel's 8086 Family manual.", "https://github.com/jpres27/8086-Instruction-Decoder");
const project4Screen = new screenState("Raycasting In a Browser", "Raycaster 2.5d fantasy RPG", "https://github.com/jpres27/Raycasting-in-a-browser");
const bioScreen = new screenState("Bio", "A really cool guy who likes CRPGs.", "");

const screens = [homeScreen, projectHome, bioScreen];
const projectScreens = [project1Screen, project2Screen, project3Screen, project4Screen];


let currentScreen = screens[0];

title.textContent = currentScreen.title;
contentSpace.textContent = currentScreen.displayText;
// github.innerHTML = "GITHUB"

displayArea.appendChild(title);
displayArea.appendChild(contentSpace);
// displayArea.appendChild(github);

ascii.classList.add('hide');
jobTitle.classList.add('hide');
title.classList.add('title');
displayArea.classList.add('text-container');
contentSpace.classList.add('text');
// github.classList.add('hide');

ascii.classList.toggle('hide');
jobTitle.classList.toggle('hide');
displayArea.classList.toggle('text-container');
contentSpace.classList.toggle('text');

window.addEventListener('keydown', (e) => {
  switch (e.key) {

    case "p": {
      currentScreen = screens[1];
      title.textContent = currentScreen.title;
      contentSpace.textContent = currentScreen.displayText;

      if(browsingProjects) {
        browsingProjects = false;
        // github.classList.toggle('hide');
      }

      if(!projectsHub) {
        projectsHub = true;
      }

      if(home){
        home = false;
        ascii.classList.toggle('hide');
        jobTitle.classList.toggle('hide');
        displayArea.classList.toggle('text-container');
        contentSpace.classList.toggle('text');
      }
    } break;

    case "b": {
      currentScreen = screens[2];
      title.textContent = currentScreen.title;
      contentSpace.textContent = currentScreen.displayText;

      if(projectsHub) {
         projectsHub = false;
      }

      if(browsingProjects) {
        browsingProjects = false;
        // github.classList.toggle('hide');
      }

      if(home){
        home = false;
        ascii.classList.toggle('hide');
        jobTitle.classList.toggle('hide');
        displayArea.classList.toggle('text-container');
        contentSpace.classList.toggle('text');
      }
    } break;

    case "h": {
      currentScreen = screens[0];
      title.textContent = currentScreen.title;
      contentSpace.textContent = currentScreen.displayText;

      if(projectsHub) {
        projectsHub = false;
      }

      if(browsingProjects) {
        browsingProjects = false;
        // github.classList.toggle('hide');
      }

      if(!home){
        home = true;
        ascii.classList.toggle('hide');
        jobTitle.classList.toggle('hide');
        displayArea.classList.toggle('text-container');
        contentSpace.classList.toggle('text');
      }
    } break; 

    case "ArrowRight": {
      if(projectsHub) {
        projectsHub = false;
        browsingProjects = true;
        // github.classList.toggle('hide');
        projectIndex = 0;
        currentScreen = projectScreens[projectIndex];
        title.textContent = currentScreen.title;
        contentSpace.textContent = currentScreen.displayText;
        // github.setAttribute('href', currentScreen.github_address);
      }

      else if(browsingProjects) {
        if(projectIndex === 3) {
          projectIndex = 0;
          currentScreen = projectScreens[projectIndex];
          title.textContent = currentScreen.title;
          contentSpace.textContent = currentScreen.displayText;
          // github.setAttribute('href', currentScreen.github_address);
        }

        else if(projectIndex < 3) {
          console.assert(currentScreen.index !== 0);
          console.assert(currentScreen.index !== 5);
          currentScreen = projectScreens[++projectIndex];
          title.textContent = currentScreen.title;
          contentSpace.textContent = currentScreen.displayText;
          // github.setAttribute('href', currentScreen.github_address);
        }
      }
    } break;

    case "ArrowLeft": {
      if(projectsHub) {
        projectsHub = false;
        browsingProjects = true;
        github.classList.toggle('hide');
        projectIndex = 3;
        currentScreen = projectScreens[projectIndex];
        title.textContent = currentScreen.title;
        contentSpace.textContent = currentScreen.displayText;
        // github.setAttribute('href', currentScreen.github_address);
      }

      else if(browsingProjects) {
        if(projectIndex === 0) {
          projectIndex = 3;
          currentScreen = projectScreens[projectIndex];
          title.textContent = currentScreen.title;
          contentSpace.textContent = currentScreen.displayText;
          // github.setAttribute('href', currentScreen.github_address);
        }

        else if(projectIndex > 0) {
          console.assert(currentScreen.index !== 5);
          currentScreen = projectScreens[--projectIndex];
          title.textContent = currentScreen.title;
          contentSpace.textContent = currentScreen.displayText;
          // github.setAttribute('href', currentScreen.github_address);
        }
      }
    } break;

    case "g": {
      if(browsingProjects) {
        window.open(currentScreen.github_address);
      }
    } break;
  }
});