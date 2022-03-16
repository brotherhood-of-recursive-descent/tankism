const DEFAULT_SIZE = 16
const DATA_SPRITE_ID = 'data-sprite-id'

const sprites = [
    '../../media/tiles/tileGrass1.png',
    '../../media/tiles/tileGrass2.png',
    '../../media/tiles/tileGrass_roadCornerLL.png',
    '../../media/tiles/tileGrass_roadCornerLR.png',
    '../../media/tiles/tileGrass_roadCornerUL.png',
    '../../media/tiles/tileGrass_roadCornerUR.png',
    '../../media/tiles/tileGrass_roadCrossing.png',
    '../../media/tiles/tileGrass_roadCrossingRound.png',
    '../../media/tiles/tileGrass_roadEast.png',
    '../../media/tiles/tileGrass_roadNorth.png',
    '../../media/tiles/tileGrass_roadSplitE.png',
    '../../media/tiles/tileGrass_roadSplitN.png',
    '../../media/tiles/tileGrass_roadSplitS.png',
    '../../media/tiles/tileGrass_roadSplitW.png',
    '../../media/tiles/tileGrass_roadTransitionE.png',
    '../../media/tiles/tileGrass_roadTransitionE_dirt.png',
    '../../media/tiles/tileGrass_roadTransitionN.png',
    '../../media/tiles/tileGrass_roadTransitionN_dirt.png',
    '../../media/tiles/tileGrass_roadTransitionS.png',
    '../../media/tiles/tileGrass_roadTransitionS_dirt.png',
    '../../media/tiles/tileGrass_roadTransitionW.png',
    '../../media/tiles/tileGrass_roadTransitionW_dirt.png',
    '../../media/tiles/tileGrass_transitionE.png',
    '../../media/tiles/tileGrass_transitionN.png',
    '../../media/tiles/tileGrass_transitionS.png',
    '../../media/tiles/tileGrass_transitionW.png',
    '../../media/tiles/tileSand1.png',
    '../../media/tiles/tileSand2.png',
    '../../media/tiles/tileSand_roadCornerLL.png',
    '../../media/tiles/tileSand_roadCornerLR.png',
    '../../media/tiles/tileSand_roadCornerUL.png',
    '../../media/tiles/tileSand_roadCornerUR.png',
    '../../media/tiles/tileSand_roadCrossing.png',
    '../../media/tiles/tileSand_roadCrossingRound.png',
    '../../media/tiles/tileSand_roadEast.png',
    '../../media/tiles/tileSand_roadNorth.png',
    '../../media/tiles/tileSand_roadSplitE.png',
    '../../media/tiles/tileSand_roadSplitN.png',
    '../../media/tiles/tileSand_roadSplitS.png',
    '../../media/tiles/tileSand_roadSplitW.png'
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
