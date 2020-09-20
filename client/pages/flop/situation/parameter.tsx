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
  Box,
} from '@material-ui/core';
import { makeStyles, createStyles, withStyles } from '@material-ui/core/styles';
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

const CustomTableCell = withStyles((theme) => ({
  root: {
    border: '1px solid rgba(224, 224, 224, 1)',
  },
  body: {
    border: '1px solid rgba(224, 224, 224, 1)',
  },
}))(TableCell);

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
            <CustomTableCell>IP Bet Freq</CustomTableCell>
            <CustomTableCell>IP Check Freq</CustomTableCell>
            <CustomTableCell>IP 33% Bet Freq</CustomTableCell>
            <CustomTableCell>IP 67% Bet Freq</CustomTableCell>
            <CustomTableCell>IP Most Common Size</CustomTableCell>
            <CustomTableCell>IP Equity</CustomTableCell>
            <CustomTableCell>OOP Bet Freq</CustomTableCell>
            <CustomTableCell>OOP Check Freq</CustomTableCell>
            <CustomTableCell>OOP 33% Bet Freq</CustomTableCell>
            <CustomTableCell>OOP 67% Bet Freq</CustomTableCell>
            <CustomTableCell>OOP Most Common Size</CustomTableCell>
            <CustomTableCell>OOP Equity</CustomTableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          <TableRow>
            <CustomTableCell>{resp.ipBetFreq}</CustomTableCell>
            <CustomTableCell>{resp.ipCheckFreq}</CustomTableCell>
            <CustomTableCell>{resp.ip33BetFreq}</CustomTableCell>
            <CustomTableCell>{resp.ip67BetFreq}</CustomTableCell>
            <CustomTableCell>
              {resp.ip33BetFreq === resp.ip67BetFreq
                ? '-'
                : resp.ip33BetFreq > resp.ip67BetFreq
                ? '33%'
                : '67%'}
            </CustomTableCell>
            <CustomTableCell>{resp.ipEquity}</CustomTableCell>
            <CustomTableCell>{resp.oopBetFreq}</CustomTableCell>
            <CustomTableCell>{resp.oopCheckFreq}</CustomTableCell>
            <CustomTableCell>{resp.oop33BetFreq}</CustomTableCell>
            <CustomTableCell>{resp.oop67BetFreq}</CustomTableCell>
            <CustomTableCell>
              {resp.oop33BetFreq === resp.oop67BetFreq
                ? '-'
                : resp.oop33BetFreq > resp.oop67BetFreq
                ? '33%'
                : '67%'}
            </CustomTableCell>
            <CustomTableCell>{resp.oopEquity}</CustomTableCell>
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
    <Grid container spacing={2}>
      {images.map((image) => {
        return (
          <Grid key={image.name} item md={4}>
            <Box my={0.5} textAlign="center">
              <Typography>{image.name}</Typography>
            </Box>
            <img src={image.url} width="100%" />
            <Box my={0.5}>
              <Typography>{image.description}</Typography>
            </Box>
          </Grid>
        );
      })}
    </Grid>
  );
};

const Parameter: Page<Props> = ({ disconnectedResp, connectedResp }: Props) => {
  const classes = useStyles({});
  return (
    <Layout className={classes.root} title="Flop parameter">
      <Container className={classes.main} component="main">
        <Box my={2}>
          <Typography variant="h4">Flop parameter</Typography>
        </Box>

        <Typography>
          Your position is
          {playerPositionTypeToString(disconnectedResp.heroPositionType)}.
        </Typography>

        <Box my={2}>
          <Typography variant="h4">Disconnected</Typography>
        </Box>

        <ParameterTable resp={disconnectedResp} />

        <Box my={2}>
          {disconnectedResp.images ? (
            <Images images={disconnectedResp.images} />
          ) : (
            <Typography variant="subtitle1">none images.</Typography>
          )}
        </Box>

        {connectedResp ? (
          <Fragment>
            <Box my={2}>
              <Typography variant="h4">Connected</Typography>
            </Box>

            <ParameterTable resp={connectedResp} />

            <Box my={2}>
              {connectedResp.images ? (
                <Images images={connectedResp.images} />
              ) : (
                <Typography variant="subtitle1">none images.</Typography>
              )}
            </Box>
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
