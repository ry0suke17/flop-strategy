import { PlayerPositionType } from '@fls-api-client/src';

export const playerPositionTypeToString = (type: PlayerPositionType) => {
  switch (type) {
    case PlayerPositionType.PlayerPositionTypeInPosition:
      return 'インポジション';
    case PlayerPositionType.PlayerPositionTypeOutOfPosition:
      return 'アウトオブポジション';
  }
};
