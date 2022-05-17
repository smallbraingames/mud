import { createWorld, createEntity, setComponent, Entity, getComponentValue } from "@mud/recs";
import { HeadlessLayer } from "../Headless";
import {
  defineStrollingComponent,
  defineLocalPositionComponent,
  defineLocalEntityTypeComponent,
  definePathComponent,
  defineMoveSpeedComponent,
  defineDestinationComponent,
  defineSelectionComponent,
  defineSelectedComponent,
  defineSelectableComponent,
  defineRockWallComponent,
} from "./components";
import {
  createDestinationSystem,
  createImpSystem,
  createPathSystem,
  createStrollingSystem,
  createSyncSystem,
  createPositionSystem,
  createSelectionSystem,
} from "./systems";
import { DEFAULT_MOVE_SPEED } from "./constants";
import { Area } from "@mud/utils";
import { createRockWallSystem } from "./systems/RockWallSystem";

/**
 * The Local layer is the thrid layer in the client architecture and extends the Headless layer.
 * Its purpose is to add components and systems for all client-only functionality, eg. strolling imps.
 */
export async function createLocalLayer(headless: HeadlessLayer) {
  const world = createWorld({ parentWorld: headless.world });

  // Components
  const LocalPosition = defineLocalPositionComponent(world);
  const LocalEntityType = defineLocalEntityTypeComponent(world);
  const Strolling = defineStrollingComponent(world);
  const Path = definePathComponent(world);
  const Destination = defineDestinationComponent(world);
  const MoveSpeed = defineMoveSpeedComponent(world);
  const Selection = defineSelectionComponent(world);
  const Selected = defineSelectedComponent(world);
  const Selectable = defineSelectableComponent(world);
  const RockWall = defineRockWallComponent(world);

  const components = {
    LocalPosition,
    LocalEntityType,
    Strolling,
    Path,
    MoveSpeed,
    Destination,
    Selection,
    Selected,
    Selectable,
    RockWall,
  };

  // Constants
  const constants = { DEFAULT_MOVE_SPEED };

  // Singleton entity
  const singletonEntity = createEntity(world);
  setComponent(Selection, singletonEntity, { x: 0, y: 0, width: 0, height: 0 });

  // API
  function selectArea(area: Area | undefined) {
    setComponent(Selection, singletonEntity, area ?? { x: 0, y: 0, width: 0, height: 0 });
  }

  function resetSelection() {
    setComponent(Selection, singletonEntity, { x: 0, y: 0, width: 0, height: 0 });
  }

  function selectEntity(entity: Entity) {
    if (getComponentValue(Selectable, entity)) setComponent(Selected, entity, {});
  }

  // Layer
  const layer = {
    world,
    components,
    parentLayers: { ...headless.parentLayers, headless },
    constants,
    api: { selectArea, selectEntity, resetSelection },
    singletonEntity,
  };

  // Systems
  createSyncSystem(layer);
  createImpSystem(layer);
  createPositionSystem(layer);
  createStrollingSystem(layer);
  createDestinationSystem(layer);
  createPathSystem(layer);
  createSelectionSystem(layer);
  createRockWallSystem(layer);

  return layer;
}