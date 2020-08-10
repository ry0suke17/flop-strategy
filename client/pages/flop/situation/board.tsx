import Layout from '@fls-components/layout';
import ConfirmationDialog from '@fls-components/confirmation_dialog';
import { Theme, Container, Typography, Button } from '@material-ui/core';
import { makeStyles, createStyles } from '@material-ui/styles';
import { Page, PageContext } from '../../_app';
import { useState, useEffect } from 'react';
import {
  PlayerPosition,
  PotType,
  HighCard,
  BoardPairType,
  BoardSuitsType,
} from '@fls-api-client/src';
import { useRouter } from 'next/router';
import { first } from '@fls-lib/strings';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {},
    main: {},
  }),
);

type Props = {
  heroPosition: PlayerPosition;
  villainPosition: PlayerPosition;
  potType: PotType;
};

const Board: Page<Props> = ({
  heroPosition,
  villainPosition,
  potType,
}: Props) => {
  const router = useRouter();

  const [errOpen, setErrOpen] = useState(false);
  const invalidPotionParam = !heroPosition || !villainPosition || !potType;

  useEffect(() => {
    if (invalidPotionParam) {
      setErrOpen(true);
    }
  }, [heroPosition, villainPosition, potType]);

  const positionParam = `hero_position=${PlayerPosition.PlayerPositionBB}&villain_position=${PlayerPosition.PlayerPositionSB}&pot_type=${PotType.PotTypeSRP}`;
  const boardParam = `high_card=${HighCard.HighCard5To7}&pair_type=${BoardPairType.BoardPairTypeUnpaired}&suits_type=${BoardSuitsType.BoardSuitsTypeMonoTone}`;
  const url = `/flop/situation/parameter?${positionParam}&${boardParam}`;

  const classes = useStyles({});

  return (
    <Layout className={classes.root} title="ボード">
      <Container className={classes.main} component="main">
        <Typography>ボード選択画面（仮）</Typography>
        <Button
          onClick={() => router.push(url)}
          variant="contained"
          color="primary"
          disabled={invalidPotionParam}
        >
          ハーカード 5~7, ペアタイプ Unpaired, スーツタイプ MonoTone を選択
        </Button>
      </Container>

      <ConfirmationDialog
        open={errOpen}
        title="エラー"
        handleClose={() => setErrOpen(false)}
        msg="パラメーターが無効です。"
      />
    </Layout>
  );
};

Board.getInitialProps = async ({ query }: PageContext) => {
  const { hero_position, villain_position, pot_type } = query;
  const heroPosition = first(hero_position) as PlayerPosition;
  const villainPosition = first(villain_position) as PlayerPosition;
  const potType = first(pot_type) as PotType;
  return {
    heroPosition,
    villainPosition,
    potType,
  };
};

export default Board;
