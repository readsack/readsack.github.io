

let currentChip = 0;

window.addEventListener("DOMContentLoaded", () => {
    let chips = []
    let tabs = document.querySelectorAll(".tab")
    tabs.forEach((_tab, i) => {
        let chip = document.createElement("div")
        chip.classList.add("chip")
        chip.setAttribute("data-tab", i)
        console.log(i)
        document.querySelector(".chips").appendChild(chip)
        chip.innerText = _tab.getAttribute("data-chip")
        chips.push(chip)
        
    });
    console.log(chips)
    let updateTabs = () => {
        tabs.forEach((tab, i) => {
            if(i == currentChip){
                tab.setAttribute("data-current", "")
            }
            else tab.removeAttribute("data-current")
        })
    }
    let updateChips = () => {

        chips.forEach(chip => {
            if (currentChip == eval(chip.attributes["data-tab"].value)) {
                chip.setAttribute("data-selected", "")
            }
            else chip.removeAttribute('data-selected')
        })
        updateTabs()
    }
    updateChips()
    //console.log(chips)
    chips.forEach(chip => {
        chip.addEventListener("click", (e) => {
            currentChip = eval(e.target.attributes["data-tab"].value)
            updateChips()
        })
    });
})



