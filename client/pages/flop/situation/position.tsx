import { Layout } from '@fls-components/layout';
import { Theme, Container } from '@material-ui/core';
import { makeStyles, createStyles } from '@material-ui/styles';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {},
    main: {},
  }),
);

const Position = () => {
  const classes = useStyles({});
  return (
    <Layout className={classes.root} title="ポジション">
      <Container className={classes.main} component="main">
        <div>ポジション選択画面（仮）</div>
      </Container>
    </Layout>
  );
};

export default Position;
