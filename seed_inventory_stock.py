import psycopg2
import random
from datetime import datetime, timedelta
from faker import Faker

fake = Faker()

DB_URL = "postgres://sunil:password123@localhost:5432/pillnet"

# ===== PROVIDED IDS =====

MEDICINE_IDS = [
    "a8c66e54-5177-4070-9694-eabc49487599",
    "7f666691-b1f9-48c4-9e4e-e86a403097df",
    "222a1d7d-1e1b-4daf-9807-678c12f1c781",
    "10c53e8e-a257-4da1-b153-6d9d5e8f6c83",
    "4bf8d268-0239-4986-b7c4-576228bdebf3",
    "1ef604cb-2fd4-4ec2-8398-264f2dcf45fa",
    "fc66e321-5d93-4b2e-8d2d-f5af65b252ec",
    "23dfe007-5265-4a5e-88c0-6d8910f2ce91",
    "116899dd-c7fc-472a-bb6b-c72e7f0cebae",
    "94d0a186-eb9a-447a-8ffb-3f3e0ce1c99c",
    "3ff8555d-cf25-424f-8343-b36a6f16af95",
    "c3049d19-9ed2-4938-9472-e4f9cb2da831",
    "349b660b-af73-431f-8303-83afd623164e",
    "d24752ea-7b87-493b-a42c-152b5eed9165",
    "bf9c6019-9eae-4539-8bf3-e624f9d9f1f5",
    "df18ad66-cd38-48e8-8545-88cc1b09ec38",
    "362e16a7-8054-4c52-bafc-c78502ebda9a",
    "36e34a3c-3c30-4394-a7f0-170ee31b5ae9",
    "639f4e60-06a9-447c-b481-96c2988716b5",
    "f46f9777-0f8c-4843-8a4a-191b9deeab7e",
    "6b298a7d-2b78-4d0b-a066-27d9c3cc3d3c",
    "d363d5dd-821f-4c9b-91da-c9ddd33621dd",
    "36b1d160-a9e4-4381-b913-b1caa568b508",
    "9c025de8-fe8e-4d69-a5bd-947e0b765c61",
    "01abbf4d-4db0-44b9-b38d-db0ff9d4bd42",
    "bb049022-c079-4d07-b061-952614177003",
    "673bde75-ac08-4ac0-9230-033618adcd70",
    "a4ec5cce-d01e-4782-8e74-427dbdf65992",
    "a66ae375-b2c2-4609-890c-a76b5ae997d2",
    "6cfdbd0f-2c6c-43f8-8cd3-7382569d659a",
    "374eaa83-0d1f-4e35-8b5f-332a2b78f1fd",
    "5a03a3fe-0af3-4763-b172-4302b49145fd",
    "0de1edcf-82fe-458f-8bd7-d750b9084df2",
    "a16ce3f3-b163-455a-a45c-a35db194427a",
    "d885770d-ceb7-44d9-99b7-61d728530f6d",
    "ad75095c-19cc-4192-b4b1-1a6b10c7e0c5",
    "c0118c7d-ee23-42a9-84e3-9e4b92c535bb",
    "d9b47c33-7cd8-4204-a81f-ac989d4d1bca",
    "455386d8-0f1a-4e25-93ce-fe276359f7cd",
]

LOCATION_IDS = [
    "d870bbee-b2be-4dca-8ef3-da3c0b3cbcb9",
    "1682faf1-c1a9-4c6a-b300-a8d84e724b25",
    "fd9ea8fe-6993-4002-88be-27d43322d806",
    "af379fa6-407f-485e-b004-f7c90001cb88",
    "ba58f370-0c51-4815-89e6-4d1d38e9ae57",
    "845e3a14-b06e-4277-af0b-813c755fa484",
    "04c98695-c494-47a2-bdc9-e77b622eb1bf",
    "8548c232-1af3-4bac-91db-eace6f6d97b5",
    "82b1c394-3629-49f8-917a-b3ab7651f1d8",
    "0cd5d144-f677-4991-a0f9-196a52971ec1",
    "5fc993ae-a399-4d4f-9c6e-1282a41c8351",
    "1366c476-d896-4a2d-990c-a3bb2d7e5c64",
    "6a255a48-e159-465a-a171-a0334d518a55",
    "cad49b9f-0db0-4d92-965b-03f839ceb145",
]

SUPPLIER_IDS = [
    "3606e956-2470-4597-82b3-00673701445d",
    "6ab1512e-d39a-4ce0-b789-ece698b4c872",
    "b51c8ec3-7d01-4a67-9cf4-282052f19548",
    "48cdced7-4918-4d1a-821f-9a4872c23cd2",
    "3d087d87-d936-459b-b90c-01c72db2f051",
    "a9b770d8-af95-4c1f-8233-fd4f1ede1506",
    "4da81280-ebaa-41b1-96c8-4e4dd632fb85",
    "6779817a-ea1d-4e30-a86c-112231b4f5a4",
    "647d1991-98b3-4163-b309-310d7e23058b",
    "849153f5-b5e0-4f76-b2a7-db898a063dca",
    "b847b2f3-195d-450f-aca0-5fa015497ed6",
    "6dafd0c5-f30d-4cb7-a9d4-40463686e8db",
    "12a7cc17-ee42-4538-b2c2-74fc6aa88bd9",
    "b82b3a9a-de9b-4b46-82c9-a3b1ef27ac41",
    "a8b5112c-9db1-41c2-9f44-6ee96a1b067e",
    "f642ba0f-dfa4-4cd6-9fda-de0f353390b9",
    "def273e0-ce40-4505-b08b-e880128c6a18",
    "c7bd6102-9ccf-4d53-8016-c1680abb6fc0",
    "718c947b-0080-4761-a084-0c2b849bec85",
    "150498ea-6b64-428b-b089-6fa7df2d7250",
    "39fcc8a3-2ee6-4222-bfdb-2d57f68f4465",
    "4b0fde17-2e06-4a26-b724-09645dbc0b3b",
]

USER_ID = "80017a2e-9ad1-4662-9348-7959b31a9bb5"
STATUSES = ["AVAILABLE", "LOW_STOCK", "RESERVED", "DAMAGED"]

# ===== DB =====
conn = psycopg2.connect(DB_URL)
cur = conn.cursor()

INSERT_SQL = """
INSERT INTO inventory_stock (
    medicine_id, batch_number,
    quantity, received_quantity, reserved_quantity, damaged_quantity,
    manufacturer_date, expiry_date, received_date,
    unit_cost_price, unit_selling_price, total_value,
    location_id, panel_code, row_number, rack_code, bin_number,
    supplier_id, status,
    stock_checked_by, stock_checked_at
)
VALUES (%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s)
"""

ROWS = 1000

for i in range(ROWS):
    received = random.randint(50, 600)
    reserved = random.randint(0, 80)
    damaged = random.randint(0, 30)
    quantity = max(received - reserved - damaged, 0)

    cost = round(random.uniform(2, 200), 2)
    sell = round(cost * random.uniform(1.2, 1.6), 2)
    total = round(quantity * cost, 2)

    mfg = datetime.now() - timedelta(days=random.randint(300, 900))
    exp = mfg + timedelta(days=random.randint(365, 900))
    recv = mfg + timedelta(days=random.randint(5, 40))

    cur.execute(
        INSERT_SQL,
        (
            random.choice(MEDICINE_IDS),
            f"BATCH-{fake.lexify(text='????').upper()}-{i}",
            quantity, received, reserved, damaged,
            mfg, exp, recv,
            cost, sell, total,
            random.choice(LOCATION_IDS),
            fake.lexify(text="P-?"),
            random.randint(0, 5),
            f"R{random.randint(1,15)}",
            random.randint(1, 1000),
            random.choice(SUPPLIER_IDS),
            random.choice(STATUSES),
            USER_ID,
            datetime.now(),
        )
    )

conn.commit()
cur.close()
conn.close()

print(f"✅ Inserted {ROWS} inventory_stock rows")
