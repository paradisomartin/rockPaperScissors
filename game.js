const piedra = {  hand:"piedra", winsAgainst:"tijera" };
const tijera = {  hand:"tijera", winsAgainst:"papel" };
const papel = {  hand:"papel", winsAgainst:"piedra" };

const handOptions = [piedra, papel, tijera]

const victory = () => {
    console.log("🚀 ~ Victory :)")
    return "Victory :)"
};

const defeat = () => {
    console.log("🚀 ~ Defeat :(")
    return "Defeat :("
};

const draw = () => {
    console.log("🚀 ~ Draw!! Try again :D")
    return "Draw!! Try again :D"
};

const randomizeComputerHand = () => {
    const randomizedValue = Math.floor(Math.random() * handOptions.length);
    return handOptions[randomizedValue]
};

const getPlayerHand = (playerInput) => {
    let playerHand = handOptions.find(hand => hand.hand === playerInput);
    // console.log("🚀 ~ getPlayerHand ~ playerHand:", playerHand)
    return playerHand ? playerHand : "Su elección no es posible, pruebe con piedra, papel o tijera"
};

const playGame = (playerInput) => {
    let playerHand = getPlayerHand(playerInput)
    console.log("🚀 Player Hand:", playerHand)
    let computerHand = randomizeComputerHand()
    console.log("🚀 Computer Hand:", computerHand)
    if (playerHand.hand === computerHand.hand) return draw()
    else if (playerHand.winsAgainst === computerHand.hand) return victory()
    else return defeat()
};

playGame("tijera")