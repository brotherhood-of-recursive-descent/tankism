const DEFAULT_SIZE = 16
const DATA_SPRITE_ID = 'data-sprite-id'

const sprites = [
    '../../assets/tiles/tileGrass1.png',
    '../../assets/tiles/tileGrass2.png',
    '../../assets/tiles/tileGrass_roadCornerLL.png',
    '../../assets/tiles/tileGrass_roadCornerLR.png',
    '../../assets/tiles/tileGrass_roadCornerUL.png',
    '../../assets/tiles/tileGrass_roadCornerUR.png',
    '../../assets/tiles/tileGrass_roadCrossing.png',
    '../../assets/tiles/tileGrass_roadCrossingRound.png',
    '../../assets/tiles/tileGrass_roadEast.png',
    '../../assets/tiles/tileGrass_roadNorth.png',
    '../../assets/tiles/tileGrass_roadSplitE.png',
    '../../assets/tiles/tileGrass_roadSplitN.png',
    '../../assets/tiles/tileGrass_roadSplitS.png',
    '../../assets/tiles/tileGrass_roadSplitW.png',
    '../../assets/tiles/tileGrass_roadTransitionE.png',
    '../../assets/tiles/tileGrass_roadTransitionE_dirt.png',
    '../../assets/tiles/tileGrass_roadTransitionN.png',
    '../../assets/tiles/tileGrass_roadTransitionN_dirt.png',
    '../../assets/tiles/tileGrass_roadTransitionS.png',
    '../../assets/tiles/tileGrass_roadTransitionS_dirt.png',
    '../../assets/tiles/tileGrass_roadTransitionW.png',
    '../../assets/tiles/tileGrass_roadTransitionW_dirt.png',
    '../../assets/tiles/tileGrass_transitionE.png',
    '../../assets/tiles/tileGrass_transitionN.png',
    '../../assets/tiles/tileGrass_transitionS.png',
    '../../assets/tiles/tileGrass_transitionW.png',
    '../../assets/tiles/tileSand1.png',
    '../../assets/tiles/tileSand2.png',
    '../../assets/tiles/tileSand_roadCornerLL.png',
    '../../assets/tiles/tileSand_roadCornerLR.png',
    '../../assets/tiles/tileSand_roadCornerUL.png',
    '../../assets/tiles/tileSand_roadCornerUR.png',
    '../../assets/tiles/tileSand_roadCrossing.png',
    '../../assets/tiles/tileSand_roadCrossingRound.png',
    '../../assets/tiles/tileSand_roadEast.png',
    '../../assets/tiles/tileSand_roadNorth.png',
    '../../assets/tiles/tileSand_roadSplitE.png',
    '../../assets/tiles/tileSand_roadSplitN.png',
    '../../assets/tiles/tileSand_roadSplitS.png',
    '../../assets/tiles/tileSand_roadSplitW.png'
]

if (!window.$) {
    new Exception("missing jquery");
}

let editor;
let exportButton;
let currentMap = [];

const generateGrid = (size) => {

    // reset
    if (editor.children().length) {
        currentMap = [];
        editor.children().remove()
    }

    for (let i = 0; i < size; i++) {

        let row = $(`<div class="row">`)
        editor.append(row)

        for (let j = 0; j < size; j++) {
            let cell = $(`<img src="" alt="" class="cell">`)

            let newIdx = 0;
            cell.attr(DATA_SPRITE_ID, newIdx);
            cell.attr('data-x', j)
            cell.attr('data-y', i)

            let tilePath = sprites[newIdx];
            cell.attr('src', tilePath)
            cell.on('click', cycleSprite)

            let mapCell = {
                file: getFileName(tilePath),
                x: j,
                y: i
            };
            currentMap.push(mapCell)
            row.append(cell)
        }
    }
}

const cycleSprite = (event) => {
    let targetElement = $(event.target);
    let spriteId = Number(targetElement.attr(DATA_SPRITE_ID))
    let cellX = Number(targetElement.attr('data-x'))
    let cellY = Number(targetElement.attr('data-y'))

    let newIdx;
    if (event.ctrlKey) {
        newIdx = spriteId === 0 ? sprites.length - 1 : spriteId - 1;
    } else {
        newIdx = spriteId === sprites.length - 1 ? 0 : spriteId + 1;
    }

    targetElement.attr(DATA_SPRITE_ID, newIdx)
    let tileName = sprites[newIdx];

    let mapCell = currentMap[cellY * DEFAULT_SIZE + cellX]
    mapCell.file = getFileName(tileName);
    mapCell.x = cellX;
    mapCell.y = cellY;

    targetElement.attr('src', tileName)
}

const registerExport = () => {
    exportButton.on('click', exportAction)
}

const exportAction = () => {

    var filename = `tankism_map_${Date.now()}.json`;

    var blob = new Blob([JSON.stringify(currentMap)], {type: 'text/plain'});
    var a = document.createElement('a');
    a.download = filename;

    a.href = window.URL.createObjectURL(blob);
    a.dataset.downloadurl = ['text/plain', a.download, a.href].join(':');
    var e = new MouseEvent('click', {"bubbles": true})
    a.dispatchEvent(e);
}


const getFileName = (fileName) => fileName.split('\\').pop().split('/').pop();


window.Mapper = {
    init: () => {
        editor = $("#editor");
        exportButton = $("#export");

        generateGrid(DEFAULT_SIZE);
        registerExport()
    }
}
