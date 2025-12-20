import psycopg2
import uuid
from datetime import datetime
from faker import Faker
import random

fake = Faker()

# Connection Config
DB_URL = "postgres://postgres:password@postgres:5432/pillnet?sslmode=disable"
# Note: Host is 'postgres' because this runs inside docker network
# If running locally, use 'localhost' and port 5433

def get_connection():
    try:
        conn = psycopg2.connect(DB_URL)
        return conn
    except:
        # Fallback for local run
        return psycopg2.connect("postgres://postgres:password@localhost:5433/pillnet?sslmode=disable")

conn = get_connection()
cur = conn.cursor()

def clean_db():
    print("🧹 Cleaning existing data...")
    tables = [
        "inventory_transactions", "inventory_stock", "alerts", 
        "medicines", "medicine_category", "manufacturers", "storage_locations"
    ]
    for t in tables:
        cur.execute(f"TRUNCATE TABLE {t} CASCADE;")
    conn.commit()

# --- SEED DATA ---

CATEGORIES = [
    "Antibiotics", "Analgesics", "Antipyretics", "Antidiabetics", 
    "Cardiovascular", "Vitamins", "Gastrointestinal"
]

MEDICINES = [
    ("Amoxicillin 500mg", "Antibiotics"),
    ("Paracetamol 650mg", "Antipyretics"),
    ("Metformin 500mg", "Antidiabetics"),
    ("Atorvastatin 10mg", "Cardiovascular"),
    ("Pantoprazole 40mg", "Gastrointestinal"),
    ("Vitamin C 500mg", "Vitamins"),
    ("Ibuprofen 400mg", "Analgesics"),
    ("Ceftriaxone Inj 1g", "Antibiotics"),
    ("Insulin Glargine", "Antidiabetics"),
    ("Aspirin 75mg", "Cardiovascular")
]

SUPPLIERS = ["Apollo Pharma", "Sun Pharma Dist", "MedPlus Vendor", "Global Health Supplies"]

def seed():
    print("🌱 Seeding Database...")
    
    # 1. Categories
    cat_map = {}
    for c in CATEGORIES:
        cat_id = str(uuid.uuid4())
        cur.execute("INSERT INTO medicine_category (id, name, description) VALUES (%s, %s, %s)", (cat_id, c, f"Category for {c}"))
        cat_map[c] = cat_id
    
    # 2. Suppliers (Manufacturers table acts as suppliers for now or do we have suppliers?)
    # Checked migration: 20251213220134_manufacturers.sql. 
    # Let's check if there is a 'suppliers' table? 
    # The existing seed script used hardcoded SUPPLIER_IDS but I don't see a `suppliers` table in the migration list I saw earlier.
    # Ah, let's look at `inventory_stock` schema in `seed_inventory_stock.py`. It has `supplier_id`.
    # Let's assume `manufacturers` table is used or there is a `users` table acting as suppliers?
    # Wait, the migration list has `manufacturers.sql`. Let's populate that and use its IDs.
    
    supp_ids = []
    for s in SUPPLIERS:
        s_id = str(uuid.uuid4())
        cur.execute("INSERT INTO manufacturers (id, name, contact_person, email, phone) VALUES (%s, %s, %s, %s, %s)", 
                   (s_id, s, fake.name(), fake.email(), fake.phone_number()[:15]))
        supp_ids.append(s_id)
        
    # 3. Storage Locations
    loc_ids = []
    for i in range(5):
        l_id = str(uuid.uuid4())
        cur.execute("INSERT INTO storage_locations (id, name, location_type, section, floor_number) VALUES (%s, %s, %s, %s, %s)",
                   (l_id, f"Shelf-{i+1}", "SHELF", "A", i))
        loc_ids.append(l_id)

    # 4. Medicines
    med_ids = []
    for name, cat in MEDICINES:
        m_id = str(uuid.uuid4())
        cur.execute("""
            INSERT INTO medicines (
                id, name, generic_name, trade_name, category_id, 
                dosage_form, strength, unit_price, mrp, 
                abc_classification, ved_classification, fsn_classification
            ) VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)
        """, (
            m_id, name, name.split()[0], name, cat_map[cat],
            "Tablet", "500mg", random.uniform(2, 50), random.uniform(5, 100),
            random.choice(['A','B','C']), random.choice(['Vital','Essential']), random.choice(['Fast','Slow'])
        ))
        med_ids.append(m_id)

    # 5. Inventory Stock
    print(f"📦 Creating stocks for {len(med_ids)} medicines...")
    for m_id in med_ids:
        # Create 2-3 batches per medicine
        for i in range(random.randint(1, 3)):
            qty = random.randint(10, 500)
            cur.execute("""
                INSERT INTO inventory_stock (
                    medicine_id, batch_number, quantity, received_quantity,
                    manufacturer_date, expiry_date, received_date, unit_cost_price, unit_selling_price,
                    location_id, supplier_id, status,
                    panel_code, row_number, rack_code, bin_number
                ) VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)
            """, (
                m_id, f"BATCH-{random.randint(1000,9999)}", qty, qty,
                datetime.now(), fake.future_date(), datetime.now(), random.uniform(5, 50), random.uniform(10, 100),
                random.choice(loc_ids), random.choice(supp_ids), "active",
                "P1", 1, "R1", 101
            ))

    conn.commit()
    print("✅ Database Seeded Successfully!")

if __name__ == "__main__":
    clean_db()
    seed()
    cur.close()
    conn.close()
