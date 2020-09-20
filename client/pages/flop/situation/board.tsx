import { ReactNode, ReactElement, Fragment, useState, useEffect } from 'react';
import Layout from '@fls-components/layout';
import ConfirmationDialog from '@fls-components/confirmation_dialog';
import {
  Theme,
  Container,
  Typography,
  Button,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
  Box,
} from '@material-ui/core';
import { makeStyles, createStyles, withStyles } from '@material-ui/core/styles';
import { Page, PageContext } from '../../_app';
import {
  PlayerPosition,
  PotType,
  HighCard,
  BoardPairType,
  BoardSuitsType,
} from '@fls-api-client/src';
import { useRouter } from 'next/router';
import { first } from '@fls-lib/strings';
import {
  boardPairTypeToString,
  boardHighCardToString,
  boardSuitsTypeToString,
} from '@fls-lib/api';
import { blue, green, red } from '@material-ui/core/colors';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {},
    main: {},
    table: {},
  }),
);

const useButtonStyle = makeStyles((theme: Theme) =>
  createStyles({
    root: (props: any) => {
      const { color, hoverColor } = props;
      return {
        color: theme.palette.getContrastText(color),
        backgroundColor: hoverColor,
        '&:hover': {
          backgroundColor: hoverColor,
        },
      };
    },
    label: {
      color: '#fff',
    },
  }),
);

type Props = {
  heroPosition: PlayerPosition;
  villainPosition: PlayerPosition;
  potType: PotType;
};

const boardPairTypes = [
  BoardPairType.BoardPairTypeUnpaired,
  BoardPairType.BoardPairTypePaired,
  BoardPairType.BoardPairTypeTrips,
];

const boardHighCardTypes = [
  HighCard.HighCardA,
  HighCard.HighCardK,
  HighCard.HighCardQ,
  HighCard.HighCardJ,
  HighCard.HighCardT,
  HighCard.HighCard8To9,
  HighCard.HighCard5To7,
  HighCard.HighCard2To4,
];

const CustomTableCell = withStyles((theme) => ({
  root: {
    border: '1px solid rgba(224, 224, 224, 1)',
  },
  body: {
    border: '1px solid rgba(224, 224, 224, 1)',
  },
}))(TableCell);

const SuitsTypeButton = ({
  heroPosition,
  villainPosition,
  potType,
  pairType,
  suitsType,
  highCard,
}: {
  heroPosition: PlayerPosition;
  villainPosition: PlayerPosition;
  potType: PotType;
  pairType: BoardPairType;
  suitsType: BoardSuitsType;
  highCard: HighCard;
}) => {
  const router = useRouter();

  const positionParam = `hero_position=${heroPosition}&villain_position=${villainPosition}&pot_type=${potType}`;
  const boardParam = `high_card=${highCard}&pair_type=${pairType}&suits_type=${suitsType}`;
  const url = `/flop/situation/parameter?${positionParam}&${boardParam}`;

  const color = (): string => {
    switch (suitsType) {
      case BoardSuitsType.BoardSuitsTypeMonoTone:
        return blue[400];
      case BoardSuitsType.BoardSuitsTypeTwoTone:
        return green[400];
      case BoardSuitsType.BoardSuitsTypeRainbow:
        return red[400];
    }
  };
  const hoverColor = (): string => {
    switch (suitsType) {
      case BoardSuitsType.BoardSuitsTypeMonoTone:
        return blue[600];
      case BoardSuitsType.BoardSuitsTypeTwoTone:
        return green[600];
      case BoardSuitsType.BoardSuitsTypeRainbow:
        return red[600];
    }
  };

  const classes = useButtonStyle({ color: color(), hoverColor: hoverColor() });

  return (
    <Box component="span" m={0.5} display="inline-block">
      <Button
        classes={{
          root: classes.root,
          label: classes.label,
        }}
        onClick={() => router.push(url)}
        variant="contained"
        color="primary"
      >
        {boardSuitsTypeToString(suitsType)}
      </Button>
    </Box>
  );
};

const SuitsTypeButtonsCell = (params: {
  heroPosition: PlayerPosition;
  villainPosition: PlayerPosition;
  potType: PotType;
  pairType: BoardPairType;
  highCard: HighCard;
}) => {
  const { highCard, pairType } = params;

  let suitsTypes: BoardSuitsType[] = [];
  switch (pairType) {
    case BoardPairType.BoardPairTypeUnpaired:
      suitsTypes.push(
        BoardSuitsType.BoardSuitsTypeMonoTone,
        BoardSuitsType.BoardSuitsTypeTwoTone,
        BoardSuitsType.BoardSuitsTypeRainbow,
      );
      break;
    case BoardPairType.BoardPairTypePaired:
      suitsTypes.push(
        BoardSuitsType.BoardSuitsTypeTwoTone,
        BoardSuitsType.BoardSuitsTypeRainbow,
      );
      break;
    case BoardPairType.BoardPairTypeTrips:
      suitsTypes.push(BoardSuitsType.BoardSuitsTypeRainbow);
      break;
  }

  return (
    <CustomTableCell align="center">
      {suitsTypes.map((suitsType) => {
        return (
          <SuitsTypeButton
            key={`${boardPairTypeToString(pairType)}-${boardHighCardToString(
              highCard,
            )}-${boardSuitsTypeToString(suitsType)}`}
            {...{ ...params, suitsType }}
          />
        );
      })}
    </CustomTableCell>
  );
};

const Board: Page<Props> = (props: Props) => {
  const router = useRouter();

  const { heroPosition, villainPosition, potType } = props;
  const [errOpen, setErrOpen] = useState(false);
  const invalidPotionParam = !heroPosition || !villainPosition || !potType;

  useEffect(() => {
    if (invalidPotionParam) {
      setErrOpen(true);
    }
  }, [heroPosition, villainPosition, potType]);

  const classes = useStyles({});

  return (
    <Layout className={classes.root} title="ボード">
      <Container className={classes.main} component="main" maxWidth="md">
        <Box my={1}>
          <Typography>Board situation</Typography>
        </Box>

        <Box my={1}>
          <TableContainer component={Paper}>
            <Table className={classes.table} aria-label="simple table">
              <TableHead>
                <TableRow>
                  <CustomTableCell></CustomTableCell>
                  {boardPairTypes.map((type) => {
                    return (
                      <CustomTableCell
                        key={boardPairTypeToString(type)}
                        align="center"
                      >
                        {boardPairTypeToString(type)}
                      </CustomTableCell>
                    );
                  })}
                </TableRow>
              </TableHead>
              <TableBody>
                {boardHighCardTypes.map((highCard) => {
                  return (
                    <TableRow key={boardHighCardToString(highCard)}>
                      <CustomTableCell align="center">
                        {boardHighCardToString(highCard)}
                      </CustomTableCell>
                      {[
                        BoardPairType.BoardPairTypeUnpaired,
                        BoardPairType.BoardPairTypePaired,
                        BoardPairType.BoardPairTypeTrips,
                      ].map((pairType) => {
                        return (
                          <SuitsTypeButtonsCell
                            key={`${boardPairTypeToString(
                              pairType,
                            )}-${boardHighCardToString(highCard)}`}
                            {...{ ...props, highCard, pairType }}
                          />
                        );
                      })}
                    </TableRow>
                  );
                })}
              </TableBody>
            </Table>
          </TableContainer>
        </Box>
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
