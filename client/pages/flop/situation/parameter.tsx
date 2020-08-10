import { Fragment } from 'react';
import Layout from '@fls-components/layout';
import {
  Theme,
  Container,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableContainer,
  TableRow,
  Paper,
  Typography,
  Grid,
} from '@material-ui/core';
import { makeStyles, createStyles } from '@material-ui/styles';
import {
  PlayerPosition,
  PotType,
  HighCard,
  BoardPairType,
  BoardSuitsType,
  BoardConnectType,
  GetFlopSituationsParameterResponse,
  GetFlopSituationsParameterResponseImages,
} from '@fls-api-client/src';
import { Page, PageContext } from '../../_app';
import { first } from '@fls-lib/strings';
import { loadAPI, playerPositionTypeToString } from '@fls-lib/api';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {},
    main: {},
  }),
);

type Props = {
  disconnectedResp: GetFlopSituationsParameterResponse;
  connectedResp?: GetFlopSituationsParameterResponse;
};

const ParameterTable = ({
  resp,
}: {
  resp: GetFlopSituationsParameterResponse;
}) => {
  return (
    <TableContainer component={Paper}>
      <Table aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>IP Bet Freq</TableCell>
            <TableCell>IP Check Freq</TableCell>
            <TableCell>IP 33% Bet Freq</TableCell>
            <TableCell>IP 67% Bet Freq</TableCell>
            <TableCell>IP Most Common Size</TableCell>
            <TableCell>IP Equity</TableCell>
            <TableCell>OOP Bet Freq</TableCell>
            <TableCell>OOP Check Freq</TableCell>
            <TableCell>OOP 33% Bet Freq</TableCell>
            <TableCell>OOP 67% Bet Freq</TableCell>
            <TableCell>OOP Most Common Size</TableCell>
            <TableCell>OOP Equity</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          <TableRow>
            <TableCell>{resp.ipBetFreq}</TableCell>
            <TableCell>{resp.ipCheckFreq}</TableCell>
            <TableCell>{resp.ip33BetFreq}</TableCell>
            <TableCell>{resp.ip67BetFreq}</TableCell>
            <TableCell>
              {resp.ip33BetFreq === resp.ip67BetFreq
                ? '-'
                : resp.ip33BetFreq > resp.ip67BetFreq
                ? '33%'
                : '67%'}
            </TableCell>
            <TableCell>{resp.ipEquity}</TableCell>
            <TableCell>{resp.oopBetFreq}</TableCell>
            <TableCell>{resp.oopCheckFreq}</TableCell>
            <TableCell>{resp.oop33BetFreq}</TableCell>
            <TableCell>{resp.oop67BetFreq}</TableCell>
            <TableCell>
              {resp.oop33BetFreq === resp.oop67BetFreq
                ? '-'
                : resp.oop33BetFreq > resp.oop67BetFreq
                ? '33%'
                : '67%'}
            </TableCell>
            <TableCell>{resp.oopEquity}</TableCell>
          </TableRow>
        </TableBody>
      </Table>
    </TableContainer>
  );
};

const Images = ({
  images,
}: {
  images: GetFlopSituationsParameterResponseImages[];
}) => {
  return (
    <Grid container>
      {images.map((image) => {
        return (
          <Grid item sm={4}>
            <img src={image.url} width="100%" />
            <Typography variant="subtitle1">{image.name}</Typography>
            <Typography variant="subtitle1">{image.description}</Typography>
          </Grid>
        );
      })}
    </Grid>
  );
};

const Parameter: Page<Props> = ({ disconnectedResp, connectedResp }: Props) => {
  const classes = useStyles({});
  return (
    <Layout className={classes.root} title="パラメータ">
      <Container className={classes.main} component="main">
        <Typography>パラメータ画面（仮）</Typography>
        <Typography variant="h3">Your Position Type</Typography>
        <Typography>
          あなたのポジションは
          {playerPositionTypeToString(disconnectedResp.heroPositionType)}
          になります。
        </Typography>
        <Typography>
          {playerPositionTypeToString(disconnectedResp.heroPositionType)}
          の項目を確認してください。
        </Typography>

        <Typography variant="h3">Disconnected Parameter</Typography>
        <ParameterTable resp={disconnectedResp} />

        {disconnectedResp.images ? (
          <Images images={disconnectedResp.images} />
        ) : (
          <Typography variant="subtitle1">画像はありません</Typography>
        )}

        {connectedResp ? (
          <Fragment>
            <Typography variant="h3">Connected Parameter</Typography>
            <ParameterTable resp={connectedResp} />
            {connectedResp.images ? (
              <Images images={connectedResp.images} />
            ) : (
              <Typography variant="subtitle1">画像はありません</Typography>
            )}
          </Fragment>
        ) : (
          ''
        )}
      </Container>
    </Layout>
  );
};

Parameter.getInitialProps = async ({ query }: PageContext) => {
  const {
    hero_position,
    villain_position,
    pot_type,
    high_card,
    pair_type,
    suits_type,
  } = query;
  const heroPosition = first(hero_position) as PlayerPosition;
  const villainPosition = first(villain_position) as PlayerPosition;
  const potType = first(pot_type) as PotType;
  const highCard = first(high_card) as HighCard;
  const boardPairType = first(pair_type) as BoardPairType;
  const boardSuitsType = first(suits_type) as BoardSuitsType;

  // Unpaired の時しかコネクトしたボードはないのでその時のみパラメータを取得する {
  var boardConnectTypes = [BoardConnectType.BoardConnectTypeDisconnect];
  if (boardPairType === BoardPairType.BoardPairTypeUnpaired) {
    boardConnectTypes.push(BoardConnectType.BoardConnectTypeConnected);
  }
  // }

  const api = loadAPI();

  const [disconnectedResp, connectedResp] = await Promise.all(
    boardConnectTypes.map((boardConnectType) =>
      api.flopStrategy.getFlopSituationsParameter({
        heroPosition,
        villainPosition,
        potType,
        highCard,
        boardPairType,
        boardSuitsType,
        boardConnectType,
      }),
    ),
  );

  return {
    disconnectedResp,
    connectedResp,
  };
};

export default Parameter;
