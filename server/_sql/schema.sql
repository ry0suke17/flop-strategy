-- card_numbers はカードの番号を表す。
CREATE TABLE card_numbers (
  -- id はカード番号の ID を表す。
  id uuid NOT NULL,
  -- display_name はカード番号の表示名を表す。
  display_name VARCHAR(1) NOT NULL CHECK(display_name <> ''),
  -- value はカード番号の値を表す。（A は 14 となる。）
  value SMALLINT NOT NULL CHECK(1 < value AND value < 15),
  -- created_at は作成時刻を表す。
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  -- updated_at は更新時刻を表す。
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id),
  UNIQUE (value)
);

-- card_suits はカードのスートを表す。
CREATE TABLE card_suits (
  -- id はカードスートの ID を表す。
  id uuid NOT NULL,
  -- display_name はカードスートの表示名を表す。
  display_name VARCHAR(1) NOT NULL CHECK(display_name <> ''),
  -- value はカードスートの値を表す。
  value SMALLINT NOT NULL CHECK(0 < value AND value < 5),
  -- created_at は作成時刻を表す。
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  -- updated_at は更新時刻を表す。
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id),
  UNIQUE (value)
);

-- cards はカードを表す。
-- cardinality: cards-card_numbers=1-1
-- cardinality: cards-card_suits=1-1
CREATE TABLE cards (
  -- id はカードの ID を表す。
  id uuid NOT NULL,
  -- card_number_id はカードスート ID を表す。
  card_number_id uuid NOT NULL,
  -- card_suit_id はカードスート ID を表す。
  card_suit_id uuid NOT NULL,
  -- created_at は作成時刻を表す。
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  -- updated_at は更新時刻を表す。
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id),
  UNIQUE (card_number_id, card_suit_id),
  FOREIGN KEY (card_number_id) REFERENCES card_numbers(id),
  FOREIGN KEY (card_suit_id) REFERENCES card_suits(id)
);

-- boards はボード（全てのコミュニティカード）を表す。
-- cardinality: boards-cards=1-3
CREATE TABLE boards (
  -- id はボードの ID を表す。
  id uuid NOT NULL,
  -- high_card_id はハイカードの id を表す。
  high_card_id uuid NOT NULL,
  -- middle_card_id はミドルカードの id を表す。
  middle_card_id uuid NOT NULL,
  -- low_card_id はローカードの id を表す。
  low_card_id uuid NOT NULL,
  -- created_at は作成時刻を表す。
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  -- updated_at は更新時刻を表す。
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id),
  UNIQUE (high_card_id, middle_card_id, low_card_id),
  FOREIGN KEY (high_card_id) REFERENCES cards(id),
  FOREIGN KEY (middle_card_id) REFERENCES cards(id),
  FOREIGN KEY (low_card_id) REFERENCES cards(id)
);

-- check_board はボードのカード順序が正しくセットされているかを確認するトリガープロシージャを表す。
CREATE FUNCTION check_board() RETURNS trigger AS $check_boards$
  DECLARE
    high_card_num_value SMALLINT;
    high_card_suit_value SMALLINT;
    middle_card_num_value SMALLINT;
    middle_card_suit_value SMALLINT;
    low_card_num_value SMALLINT;
    low_card_suit_value SMALLINT;
  BEGIN
    SELECT high_card_num.value, high_card_suit.value, middle_card_num.value, middle_card_suit.value, low_card_num.value, low_card_suit.value
      INTO high_card_num_value, high_card_suit_value, middle_card_num_value, middle_card_suit_value, low_card_num_value, low_card_suit_value
      FROM (SELECT NEW.*) board
      INNER JOIN cards AS high_card ON board.high_card_id = high_card.id
      INNER JOIN card_numbers AS high_card_num ON high_card.card_number_id = high_card_num.id
      INNER JOIN card_suits AS high_card_suit ON high_card.card_suit_id = high_card_suit.id
      INNER JOIN cards AS middle_card ON board.middle_card_id = middle_card.id
      INNER JOIN card_numbers AS middle_card_num ON middle_card.card_number_id = middle_card_num.id
      INNER JOIN card_suits AS middle_card_suit ON middle_card.card_suit_id = middle_card_suit.id
      INNER JOIN cards AS low_card ON board.low_card_id = low_card.id
      INNER JOIN card_numbers AS low_card_num ON low_card.card_number_id = low_card_num.id
      INNER JOIN card_suits AS low_card_suit ON low_card.card_suit_id = low_card_suit.id;

    IF (high_card_num_value < middle_card_num_value) THEN
      RAISE EXCEPTION 'The card number order is invalid. The middle card should be smaller or equal than the high card. high_card_num_value: %, middle_card_num_value: %', high_card_num_value, middle_card_num_value;
    END IF;
    IF (high_card_num_value = middle_card_num_value AND high_card_suit_value < middle_card_suit_value) THEN
      RAISE EXCEPTION 'The card suit order is invalid. The middle card should be smaller or equal than the high card. high_card_suit_value: %, middle_card_suit_value: %', high_card_suit_value, middle_card_suit_value;
    END IF;

    IF (middle_card_num_value < low_card_num_value) THEN
      RAISE EXCEPTION 'The card number order is invalid. The low card should be smaller or equal than the middle card. middle_card_num_value: %, low_card_num_value: %', middle_card_num_value, low_card_num_value;
    END IF;
    IF (middle_card_num_value = low_card_num_value AND middle_card_suit_value < low_card_suit_value) THEN
      RAISE EXCEPTION 'The card suit order is invalid. The low card should be smaller or equal than the middle card. middle_card_suit_value: %, low_card_suit_value: %', middle_card_suit_value, low_card_suit_value;
    END IF;

    RETURN NEW;
  END;
$check_boards$ LANGUAGE plpgsql;

-- check_boards はボードのカード順序が正しくセットされているかを確認するトリガーを表す。
CREATE TRIGGER check_board BEFORE INSERT OR UPDATE ON boards
  FOR EACH ROW EXECUTE PROCEDURE check_board();

-- player_positions はプレイヤーのポジションを表す。
CREATE TABLE player_positions (
  -- id はポジションの ID を表す。
  id uuid NOT NULL,
  -- display_name はポジションの表示名を表す。
  display_name VARCHAR(3) NOT NULL CHECK(display_name <> ''),
  -- post_flop_action_order はポストフロップでのアクションの順番を表す。
  post_flop_action_order SMALLINT NOT NULL CHECK(0 < post_flop_action_order AND post_flop_action_order < 7), -- 6 MAX のみを想定している。
  -- created_at は作成時刻を表す。
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  -- updated_at は更新時刻を表す。
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id),
  UNIQUE (post_flop_action_order)
);

-- pot_type はポットタイプを表す。
CREATE TYPE pot_type AS ENUM (
    -- シングルレイズドポットを表す。
    'SRP',
    -- 3 ベットポットを表す。
    '3BET',
    -- 4 ベットポットを表す。
    '4BET'
);

-- heads_up_situations はヘッズアップのシチュエーションを表す。
-- cardinality: heads_up_situations-player_positions=1-2
CREATE TABLE heads_up_situations (
  -- id はヘッズアップシチュエーションの ID を表す。
  id uuid NOT NULL,
  -- in_position_id はインポジションであるポジション ID を表す。
  in_position_id uuid NOT NULL,
  -- out_of_position_id はアウトオブポジションであるポジション ID を表す。
  out_of_position_id uuid NOT NULL,
  -- pot_type はポットタイプを表す。
  pot_type pot_type NOT NULL,
  -- created_at は作成時刻を表す。
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  -- updated_at は更新時刻を表す。
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id),
  UNIQUE (in_position_id, out_of_position_id),
  FOREIGN KEY (in_position_id) REFERENCES player_positions(id),
  FOREIGN KEY (out_of_position_id) REFERENCES player_positions(id)
);

-- check_heads_up_situation はヘッズアップシチュエーションのプレイヤーポジション順序が正しくセットされているかを確認するトリガープロシージャを表す。
CREATE FUNCTION check_heads_up_situation() RETURNS trigger AS $check_heads_up_situation$
  DECLARE
    ip_order SMALLINT;
    oop_order SMALLINT;
  BEGIN
    SELECT ip.post_flop_action_order, oop.post_flop_action_order
      INTO ip_order, oop_order
      FROM (SELECT NEW.*) heads_up
      INNER JOIN player_positions AS ip ON heads_up.in_position_id = ip.id
      INNER JOIN player_positions AS oop ON heads_up.out_of_position_id = oop.id;

    IF (ip_order < oop_order) THEN
      RAISE EXCEPTION 'The player position order is invalid. OOP player must come before IP player. ip_order: %, oop_order: %', ip_order, oop_order;
    END IF;

    RETURN NEW;
  END;
$check_heads_up_situation$ LANGUAGE plpgsql;

-- check_heads_up_situation はヘッズアップシチュエーションのプレイヤーポジション順序が正しくセットされているかを確認するトリガーを表す。
CREATE TRIGGER check_heads_up_situation BEFORE INSERT OR UPDATE ON heads_up_situations
  FOR EACH ROW EXECUTE PROCEDURE check_heads_up_situation();

-- flop_situations はフロップのシチュエーションを表す。
-- cardinality: flop_situations-boards=1-1
-- cardinality: flop_situations-player_positions=1-2
CREATE TABLE flop_situations (
  -- id はフロップシチューションの ID を表す。
  id uuid NOT NULL,
  -- board_id はボード ID を表す。
  board_id uuid NOT NULL,
  -- heads_up_situation_id はヘッズアップシチュエーションの ID を表す。
  heads_up_situation_id uuid NOT NULL,
  -- created_at は作成時刻を表す。
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  -- updated_at は更新時刻を表す。
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id),
  UNIQUE (board_id, heads_up_situation_id),
  FOREIGN KEY (board_id) REFERENCES boards(id),
  FOREIGN KEY (heads_up_situation_id) REFERENCES heads_up_situations(id)
);

-- flop_situation_parameters はフロップシチュエーションのパラメータを表す。
-- cardinality: flop_situations-flop_situation_parameters=1-1
CREATE TABLE flop_situation_parameters (
  -- flop_situations_id はフロップシチュエーションの ID を表す。
  flop_situations_id uuid NOT NULL,
  -- out_of_position_bet_frequency はアウトオブポジションのベット頻度を表す。
  out_of_position_bet_frequency real NOT NULL,
  -- in_position_bet_frequency はインポジションのベット頻度を表す。
  in_position_bet_frequency real NOT NULL,
  -- out_of_position_check_frequency はアウトオブポジションのチェック頻度を表す。
  out_of_position_check_frequency real NOT NULL,
  -- in_position_check_frequency はインポジションのチェック頻度を表す。
  in_position_check_frequency real NOT NULL,
  -- out_of_position_equity はアウトオブポジションのエクイティを表す。
  out_of_position_equity real NOT NULL,
  -- in_position_equity はインポジションのエクイティを表す。
  in_position_equity real NOT NULL,
  -- out_of_position_33_bet_frequency はアウトオブポジションの 33% ベットの頻度を表す。
  out_of_position_33_bet_frequency real NOT NULL,
  -- in_position_33_bet_frequency はアウトオブポジションの 33% ベットの頻度を表す。
  in_position_33_bet_frequency real NOT NULL,
  -- out_of_position_67_bet_frequency はアウトオブポジションの 33% ベットの頻度を表す。
  out_of_position_67_bet_frequency real NOT NULL,
  -- in_position_67_bet_frequency はアウトオブポジションの 67% ベットの頻度を表す。
  in_position_67_bet_frequency real NOT NULL,
  -- created_at は作成時刻を表す。
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  -- updated_at は更新時刻を表す。
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  PRIMARY KEY (flop_situations_id),
  FOREIGN KEY (flop_situations_id) REFERENCES flop_situations(id)
);

-- flop_situation_images はフロップシチュエーションの画像を表す。
-- cardinality: flop_situations-flop_situation_images=1-1
CREATE TABLE flop_situation_images (
  -- flop_situations_id はフロップシチュエーションの ID を表す。
  flop_situations_id uuid NOT NULL,
  -- url は画像の URL を表す。
  url VARCHAR(1000) NOT NULL CHECK(url <> ''),
  -- description は画像の説明文を表す。（説明文がない場合は空文字が入る想定。）
  description VARCHAR(3000) NOT NULL,
  -- created_at は作成時刻を表す。
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  -- updated_at は更新時刻を表す。
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  PRIMARY KEY (flop_situations_id),
  FOREIGN KEY (flop_situations_id) REFERENCES flop_situations(id)
);