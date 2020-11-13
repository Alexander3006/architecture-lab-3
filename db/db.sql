CREATE TABLE "Machines"(
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR,
    "cpuCount" smallint check ( "cpuCount" > 0 )
);

CREATE TABLE "Disks" (
    "id" SERIAL PRIMARY KEY,
    "space" INTEGER CHECK ("space" > 0)
);

CREATE TABLE "MachineComponents" (
  "id" SERIAL PRIMARY KEY,
  "machineId" INTEGER,
  "diskId" INTEGER UNIQUE,
  FOREIGN KEY ("machineId") REFERENCES "Machines"("id") ON DELETE CASCADE,
  FOREIGN KEY ("diskId") REFERENCES "Disks"("id") ON DELETE CASCADE
);

CREATE OR REPLACE FUNCTION "MachinesInfo"()
returns TABLE (
    id int,
    "name" varchar,
    "cpuCount" smallint,
    "totalDiskSpace" bigint
)
language plpgsql
as $code$
	BEGIN
        RETURN QUERY SELECT M.id, M.name, M."cpuCount", sum(D.space) as "totalDiskSpace"
        FROM "Machines" as M
        JOIN "MachineComponents" as MC on M.id = MC."machineId"
        JOIN "Disks" as D on D.id = MC."diskId"
        GROUP BY M.id;
	END
$code$;