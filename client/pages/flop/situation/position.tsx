import Layout from '@fls-components/layout';
import { Theme, Container, Typography, Button } from '@material-ui/core';
import { makeStyles, createStyles } from '@material-ui/styles';
import { useRouter } from 'next/router';
import { PlayerPosition, PotType } from '@fls-api-client/src';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {},
    main: {},
  }),
);

const Position = () => {
  const router = useRouter();

  const positionParam = `hero_position=${PlayerPosition.PlayerPositionBB}&villain_position=${PlayerPosition.PlayerPositionSB}&pot_type=${PotType.PotTypeSRP}`;
  const url = `/flop/situation/board?${positionParam}`;

  const classes = useStyles({});

  return (
    <Layout className={classes.root} title="ポジション">
      <Container className={classes.main} component="main">
        <Typography>ポジション選択画面（仮）</Typography>
        <Button
          onClick={() => router.push(url)}
          variant="contained"
          color="primary"
        >
          Hero BB, Villain SB, ポットタイプ SRP を選択
        </Button>
      </Container>
    </Layout>
  );
};

export default Position;
