import {
  PlayerPositionType,
  BoardPairType,
  HighCard,
  BoardSuitsType,
} from '@fls-api-client/src';

export const playerPositionTypeToString = (type: PlayerPositionType) => {
  switch (type) {
    case PlayerPositionType.PlayerPositionTypeInPosition:
      return 'インポジション';
    case PlayerPositionType.PlayerPositionTypeOutOfPosition:
      return 'アウトオブポジション';
  }
};

export const boardPairTypeToString = (type: BoardPairType) => {
  switch (type) {
    case BoardPairType.BoardPairTypeUnpaired:
      return 'Unpaired';
    case BoardPairType.BoardPairTypePaired:
      return 'Paired';
    case BoardPairType.BoardPairTypeTrips:
      return 'trips';
  }
};

export const boardHighCardToString = (highCard: HighCard) => {
  switch (highCard) {
    case HighCard.HighCardA:
      return 'A';
    case HighCard.HighCardK:
      return 'K';
    case HighCard.HighCardQ:
      return 'Q';
    case HighCard.HighCardJ:
      return 'J';
    case HighCard.HighCardT:
      return 'T';
    case HighCard.HighCard8To9:
      return '8-9';
    case HighCard.HighCard5To7:
      return '5-7';
    case HighCard.HighCard2To4:
      return '2-4';
  }
};

export const boardSuitsTypeToString = (suitsType: BoardSuitsType) => {
  switch (suitsType) {
    case BoardSuitsType.BoardSuitsTypeMonoTone:
      return 'mono';
    case BoardSuitsType.BoardSuitsTypeTwoTone:
      return 'tt';
    case BoardSuitsType.BoardSuitsTypeRainbow:
      return 'r';
  }
};
